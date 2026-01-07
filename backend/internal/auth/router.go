package auth

import (
	gobetterauth "github.com/GoBetterAuth/go-better-auth"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(app *fiber.App, auth *gobetterauth.Auth) {
	app.Use(adaptor.HTTPHandler(auth.Handler()))
}
