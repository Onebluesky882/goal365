package auth

import gobetterauth "github.com/GoBetterAuth/go-better-auth"

type Auth struct {
	Auth *gobetterauth.Auth
}

func New() *Auth {
	cfg := NewButterAuthConfig()
	a := gobetterauth.New(cfg)

	// a.RunMigrations()
	return &Auth{
		Auth: a,
	}
}
