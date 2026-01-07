package auth

import (
	"os"

	gobetterauthconfig "github.com/GoBetterAuth/go-better-auth/config"
	gobetterauthmodels "github.com/GoBetterAuth/go-better-auth/models"
)

func NewButterAuthConfig() *gobetterauthmodels.Config {
	return gobetterauthconfig.NewConfig(
		gobetterauthconfig.WithBaseURL("http://localhost:3000"),
		gobetterauthconfig.WithBasePath("/api/auth"),
		gobetterauthconfig.WithDatabase(gobetterauthmodels.DatabaseConfig{
			Provider: "postgres",
			URL:      os.Getenv("DATABASE_URL"),
		}),
		gobetterauthconfig.WithEmailPassword(gobetterauthmodels.EmailPasswordConfig{
			Enabled: true,
		}),
	)
}
