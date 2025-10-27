package auth

import (
	"encontradev/internal/repository"
	"os"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Auth struct {
	Secret            []byte
	Repository        *repository.Repository
	GoogleOauthConfig *oauth2.Config
	GithubOauthConfig *oauth2.Config
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func SetAuth(repository *repository.Repository) (auth *Auth, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	secret := os.Getenv("JWT_SECRET")
	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	googleRedirectUrl := os.Getenv("GOOGLE_REDIRECT_URL")

	var googleOauthConfig = &oauth2.Config{
		ClientID:     googleClientId,
		ClientSecret: googleClientSecret,
		RedirectURL:  googleRedirectUrl,
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}

	a := &Auth{
		Secret:            []byte(secret),
		GoogleOauthConfig: googleOauthConfig,
		Repository:        repository,
	}
	return a, err
}
