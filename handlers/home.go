package handlers

import (
	home "goth-template/views"
	"net/http"
)

func (h *Handler) HandleHome(w http.ResponseWriter, r *http.Request) error {
	u, err := h.auth.GetUserSession(r)
	if err != nil {
		return err
	}
	return Render(w, r, home.Index(u))
}

func HandleFoo(w http.ResponseWriter, r *http.Request) error {
	_, err := w.Write([]byte("Welcome foo"))
	return err
}
