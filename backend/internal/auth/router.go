package auth

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type VerifyLineTokenRequest struct {
	IDToken       string `json:"idToken"`
	UserID        string `json:"userId"`
	DisplayName   string `json:"displayName"`
	PictureURL    string `json:"pictureUrl"`
	StatusMessage string `json:"statusMessage"`
}

func RegisterRoute(app *fiber.App) {
	api := app.Group("/api")

	// New endpoint: Verify ID token sent from frontend
	api.Post("/auth/verify-line", func(c *fiber.Ctx) error {
		var req VerifyLineTokenRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		// Verify the ID token
		user, err := verifyLineIDToken(req.IDToken)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{
				"error":   "Invalid ID token",
				"details": err.Error(),
			})
		}

		fmt.Printf("✅ LINE user verified: %s (%s)\n", user.Name, user.Sub)

		// TODO: Create or update user in your database
		// db.CreateOrUpdateUser(user)

		// TODO: Generate your own JWT/session token
		// sessionToken := generateSessionToken(user.Sub)

		return c.JSON(fiber.Map{
			"success": true,
			"user": fiber.Map{
				"id":      user.Sub,
				"name":    user.Name,
				"picture": user.Picture,
				"email":   user.Email,
			},
			// "token": sessionToken, // Your backend session token
		})
	})

	// Keep the callback endpoint for direct web login (without LIFF)
	// This won't work with LIFF PKCE but is useful for LINE Login button
	app.Get("/", handleLineCallback)
	api.Get("/auth/callback/line", handleLineCallback)
}

func handleLineCallback(c *fiber.Ctx) error {
	code := c.Query("code")

	if code == "" {
		return c.SendString("Welcome to Goal65 API - Server is running!")
	}

	// This endpoint is for non-LIFF LINE Login
	// For LIFF, use the /api/auth/verify-line endpoint instead
	return c.JSON(fiber.Map{
		"message": "For LIFF login, please use the LIFF SDK on the frontend and send the ID token to /api/auth/verify-line",
	})
}
