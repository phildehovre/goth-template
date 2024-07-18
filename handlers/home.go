package handlers

import (
	home "goth-template/views"
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, home.Index())
}

func HandleFoo(w http.ResponseWriter, r *http.Request) error {
	_, err := w.Write([]byte("Welcome foo"))
	return err
}
