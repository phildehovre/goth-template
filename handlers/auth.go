package handlers

import (
	"context"
	home "goth-template/views"
	auth "goth-template/views/auth"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/markbates/goth/gothic"
)

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) error {

	return Render(w, r, auth.Login())
}

func appendProviderStringToContext(r *http.Request) *http.Request {
	provStr := chi.URLParam(r, "provider")
	ctx := context.WithValue(r.Context(), "provider", provStr)
	return r.WithContext(ctx)
}

func (h *Handler) HandleProviderLogin(w http.ResponseWriter, r *http.Request) error {

	if u, err := gothic.CompleteUserAuth(w, r); err == nil {
		slog.Info("User already authenticated: %v", u)
		Render(w, r, home.Index())
	} else {
		r = appendProviderStringToContext(r)
		gothic.BeginAuthHandler(w, r)

	}
	return nil
}

func (h *Handler) HandlerAuthCallbackFunc(w http.ResponseWriter, r *http.Request) error {
	r = appendProviderStringToContext(r)
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		return err
	}

	if err := h.auth.StoreUserSession(w, r, user); err != nil {
		return err
	}

	w.Header().Set("Location", "/foo")
	w.WriteHeader(http.StatusTemporaryRedirect)
	return nil

}
