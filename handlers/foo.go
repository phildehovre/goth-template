package handlers

import (
	foo "goth-template/views"
	"net/http"
)

func HandleFoo(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, foo.Index())
}
