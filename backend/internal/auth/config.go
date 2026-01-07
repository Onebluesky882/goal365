package auth

import (
	"os"

	gobetterauthconfig "github.com/GoBetterAuth/go-better-auth/config"
	gobetterauthmodels "github.com/GoBetterAuth/go-better-auth/models"
)

func NewButterAuthConfig() *gobetterauthmodels.Config {

	return gobetterauthconfig.NewConfig(
		gobetterauthconfig.WithMode(gobetterauthmodels.ModeStandalone),
		gobetterauthconfig.WithDatabase(gobetterauthmodels.DatabaseConfig{
			Provider: "postgres",
			URL:      os.Getenv("DATABASE_URL"),
		}),
		gobetterauthconfig.WithBaseURL("https://your-domain.com"),
		gobetterauthconfig.WithEmailPassword(gobetterauthmodels.EmailPasswordConfig{
			Enabled: true,
		}),
		lineProvider(),
	)

}
