package auth

import (
	"fmt"
	"goth-template/config"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

type AuthService struct {
	provider string
}

func NewAuthService(store *sessions.CookieStore) *AuthService {
	gothic.Store = store

	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_CLIENT_ID"),
			os.Getenv("GOOGLE_SECRET"),
			buildCallbackURL("google"),
		),
	)

	return &AuthService{}
}

const SessionName = "session-name"

func (s *AuthService) StoreUserSession(w http.ResponseWriter, r *http.Request, user goth.User) error {
	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.
	session, _ := gothic.Store.Get(r, SessionName)
	// Set some session values.
	session.Values["user"] = user
	// Save it before we write to the response/return from the handler.
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	return nil
}

func (s *AuthService) GetUserSession(r *http.Request) (goth.User, error) {
	session, err := gothic.Store.Get(r, SessionName)
	if err != nil {
		return goth.User{}, err
	}
	u := session.Values["user"]
	if u == nil {
		return goth.User{}, fmt.Errorf("User is not authenticated %v", u)

	}
	return u.(goth.User), nil
}

func RequireAuth(f http.HandlerFunc, auth *AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := auth.GetUserSession(r)

		if err != nil {
			log.Println("User is not authenticated")
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		log.Printf("user is authenticated! user: %v", session.Name)

		f(w, r)
	}

}

func buildCallbackURL(provider string) string {
	return fmt.Sprintf("%s:%s/auth/%s/callback", config.Envs.PublicHost, config.Envs.Port, provider)
}
