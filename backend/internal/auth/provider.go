package auth

import (
	"os"

	gobetterauthmodels "github.com/GoBetterAuth/go-better-auth/models"
)

func lineProvider() gobetterauthmodels.SocialProvidersConfig {
	return gobetterauthmodels.SocialProvidersConfig{
		"line": gobetterauthmodels.OAuth2ProviderConfig{
			Enabled:      true,
			ClientID:     os.Getenv("LINE_CLIENT_ID"),
			ClientSecret: os.Getenv("LINE_CLIENT_SECRET"),
			RedirectURL:  "http://localhost:3000/auth/callback/line",
			Scopes: []string{
				"profile",
				"openid",
				"email",
			},
		},
	}
}
