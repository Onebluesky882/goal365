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
				RedirectURL:  "https://onebluesky882.jprq.live/auth/callback/line",

				AuthURL:     "https://access.line.me/oauth2/v2.1/authorize",
				TokenURL:    "https://api.line.me/oauth2/v2.1/token",
				UserInfoURL: "https://api.line.me/v2/profile",

				Scopes: []string{"profile", "openid", "email"},
			},
		},
	)
}
