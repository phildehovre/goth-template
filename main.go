package main

import (
	"log"
	"net/http"
	"os"

	"goth-template/handlers"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	r := chi.NewMux()

	r.Handle("/*", public())

	r.Get("/", handlers.Make(handlers.HandleHome))

	r.Get("/auth/{provider}", handlers.Make(nil))
	r.Get("/auth/{provider}/callback", handlers.Make(nil))
	r.Get("/auth/logout/{provider}", handlers.Make(nil))
	r.Get("/auth/login", handlers.Make(handlers.HandleLogin))

	listenAddr := os.Getenv("LISTEN_ADDR")
	http.ListenAndServe(listenAddr, r)
}
