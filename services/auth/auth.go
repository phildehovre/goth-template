package auth

import (
	"fmt"
	"goth-template/config"
	"os"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/twitter"
)

type AuthService struct {
	provider string
}

func NewAuthService() *AuthService {
	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_CLIENT_KEY"), os.Getenv("GOOGLE_SECRET"), buildCallbackURL("google")),
		facebook.New(os.Getenv("FACEBOOK_CLIENT_KEY"), os.Getenv("FACEBOOK_SECRET"), buildCallbackURL("google")),
		twitter.New(os.Getenv("TWITTER_CLIENT_KEY"), os.Getenv("TWITTER_SECRET"), buildCallbackURL("google")),
	)

	return &AuthService{}
}

func buildCallbackURL(provider string) string {
	return fmt.Sprintf("%s:%s/auth/%s/callback", config.Envs.PublicHost, config.Envs, provider)
}
