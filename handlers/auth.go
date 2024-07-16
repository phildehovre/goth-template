package handlers

import (
	"fmt"
	home "goth-template/views"
	"goth-template/views/auth"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/markbates/goth/gothic"
)

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) error {

	return Render(w, r, auth.Login())
}

func (h *Handler) HandleProviderLogin(w http.ResponseWriter, r *http.Request) error {
	provider := chi.URLParam(r, "provider")
	fmt.Println(provider)
	var err error
	if u, err := gothic.CompleteUserAuth(w, r); err == nil {
		slog.Info("User already authenticated: %v", u)
		Render(w, r, home.Index())
	} else {
		gothic.BeginAuthHandler(w, r)

	}
	return err
}

func HandlerAuthCallbackFunc(w http.ResponseWriter, r *http.Request) error {

	provider := chi.URLParam(r, "provider")
	fmt.Println(provider)
	_, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		return err
	}
	return nil

}
