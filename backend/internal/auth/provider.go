package auth

import (
	"os"

	gobetterauthconfig "github.com/GoBetterAuth/go-better-auth/config"
	gobetterauthmodels "github.com/GoBetterAuth/go-better-auth/models"
)

func lineProvider() gobetterauthmodels.ConfigOption {
	return gobetterauthconfig.WithSocialProviders(
		gobetterauthmodels.SocialProvidersConfig{
			"line": gobetterauthmodels.OAuth2ProviderConfig{
				Enabled:      true,
				ClientID:     os.Getenv("LINE_CLIENT_ID"),
				ClientSecret: os.Getenv("LINE_CLIENT_SECRET"),

				Scopes: []string{"profile", "openid", "email"},
			},
		},
	)
}
