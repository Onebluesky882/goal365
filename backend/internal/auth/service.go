package auth

import (
	"bytes"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

// ============================================================================
// TYPE DEFINITIONS
// ============================================================================

// Response ของ LINE token endpoint
type LineTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	IDToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

// ข้อมูล user จาก LINE id_token
type LineUser struct {
	Sub           string `json:"sub"` // LINE userId
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	Email         string `json:"email,omitempty"`
	EmailVerified bool   `json:"email_verified,omitempty"`
}

// LINE JWK (JSON Web Key) structure
type LineJWK struct {
	Keys []struct {
		Kty string `json:"kty"`
		Use string `json:"use"`
		Kid string `json:"kid"`
		N   string `json:"n"`
		E   string `json:"e"`
		Alg string `json:"alg"`
	} `json:"keys"`
}

// ============================================================================
// TOKEN EXCHANGE
// ============================================================================

// แลก code → access_token + id_token
func lineExchangeCode(code string, liffRedirectUri string, codeVerifier string) (*LineTokenResponse, error) {
	clientID := os.Getenv("LINE_CLIENT_ID")
	clientSecret := os.Getenv("LINE_CLIENT_SECRET")
	
	// Use the liffRedirectUri from the callback if provided
	// LINE sends this as the actual redirect_uri it used
	redirectURI := liffRedirectUri
	if redirectURI == "" {
		redirectURI = os.Getenv("LINE_REDIRECT_URI")
	}

	fmt.Printf("Using redirect_uri: %s\n", redirectURI)

	// Build form data - use url.Values to properly encode
	formData := url.Values{}
	formData.Set("grant_type", "authorization_code")
	formData.Set("code", code)
	formData.Set("redirect_uri", redirectURI)
	formData.Set("client_id", clientID)
	formData.Set("client_secret", clientSecret)
	
	// Add code_verifier for PKCE flow (if provided)
	if codeVerifier != "" {
		formData.Set("code_verifier", codeVerifier)
		fmt.Printf("Using PKCE with code_verifier\n")
	}

	fmt.Printf("Form data: %s\n", formData.Encode())

	req, err := http.NewRequest("POST", "https://api.line.me/oauth2/v2.1/token", 
		bytes.NewBufferString(formData.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("line token error (status %d): %s", resp.StatusCode, string(body))
	}

	var tokenResp LineTokenResponse
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return nil, err
	}

	return &tokenResp, nil
}

// ============================================================================
// ID TOKEN VERIFICATION
// ============================================================================

// ตรวจสอบ JWT id_token ของ LINE
func verifyLineIDToken(idToken string) (*LineUser, error) {
	// Parse token without verification first to get kid from header
	token, _, err := new(jwt.Parser).ParseUnverified(idToken, jwt.MapClaims{})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	// Get kid from header
	kid, ok := token.Header["kid"].(string)
	if !ok {
		return nil, errors.New("missing kid in JWT header")
	}

	// Fetch LINE's public keys
	resp, err := http.Get("https://api.line.me/oauth2/v2.1/certs")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch LINE certs: %w", err)
	}
	defer resp.Body.Close()

	var jwks LineJWK
	if err := json.NewDecoder(resp.Body).Decode(&jwks); err != nil {
		return nil, fmt.Errorf("failed to decode JWKs: %w", err)
	}

	// Find the matching key
	var matchedKey *struct {
		Kty string `json:"kty"`
		Use string `json:"use"`
		Kid string `json:"kid"`
		N   string `json:"n"`
		E   string `json:"e"`
		Alg string `json:"alg"`
	}

	for i := range jwks.Keys {
		if jwks.Keys[i].Kid == kid {
			matchedKey = &jwks.Keys[i]
			break
		}
	}

	if matchedKey == nil {
		return nil, errors.New("no matching LINE public key found")
	}

	// Convert JWK to RSA public key
	pubKey, err := jwkToRSAPublicKey(matchedKey.N, matchedKey.E)
	if err != nil {
		return nil, fmt.Errorf("failed to convert JWK to RSA public key: %w", err)
	}

	// Verify and parse the token with the public key
	parsedToken, err := jwt.Parse(idToken, func(token *jwt.Token) (interface{}, error) {
		// Verify the signing method
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return pubKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to verify token: %w", err)
	}

	if !parsedToken.Valid {
		return nil, errors.New("invalid id_token")
	}

	// Extract claims
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("cannot parse claims")
	}

	// Build LineUser from claims
	user := &LineUser{
		Sub:     getStringClaim(claims, "sub"),
		Name:    getStringClaim(claims, "name"),
		Picture: getStringClaim(claims, "picture"),
		Email:   getStringClaim(claims, "email"),
	}

	if ev, ok := claims["email_verified"].(bool); ok {
		user.EmailVerified = ev
	}

	return user, nil
}

// ============================================================================
// HELPER FUNCTIONS
// ============================================================================

// Helper function to convert JWK to RSA Public Key
func jwkToRSAPublicKey(nStr, eStr string) (*rsa.PublicKey, error) {
	// Decode base64url-encoded n (modulus)
	nBytes, err := base64.RawURLEncoding.DecodeString(nStr)
	if err != nil {
		return nil, fmt.Errorf("failed to decode n: %w", err)
	}

	// Decode base64url-encoded e (exponent)
	eBytes, err := base64.RawURLEncoding.DecodeString(eStr)
	if err != nil {
		return nil, fmt.Errorf("failed to decode e: %w", err)
	}

	// Convert bytes to big.Int
	n := new(big.Int).SetBytes(nBytes)
	
	// Convert exponent bytes to int
	e := 0
	for _, b := range eBytes {
		e = e<<8 | int(b)
	}

	// Create RSA public key
	pubKey := &rsa.PublicKey{
		N: n,
		E: e,
	}

	return pubKey, nil
}

// Helper function to safely get string claims
func getStringClaim(claims jwt.MapClaims, key string) string {
	if val, ok := claims[key].(string); ok {
		return val
	}
	return ""
}