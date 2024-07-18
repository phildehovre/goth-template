package main

import (
	"log"
	"net/http"
	"os"

	"goth-template/config"
	"goth-template/handlers"
	"goth-template/services/auth"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	r := chi.NewMux()

	sessionsStore := auth.NewCookieStore(auth.SessionsOptions{
		CookiesKey: config.Envs.CookiesAuthSecret,
		MaxAge:     config.Envs.CookiesAuthAgeInSeconds,
		Secure:     config.Envs.CookiesAuthIsSecure,
		HttpOnly:   config.Envs.CookiesAuthIsHttpOnly,
	})
	authProvider := auth.NewAuthService(sessionsStore)
	handler := handlers.New(nil, authProvider)

	r.Handle("/*", public())

	r.Get("/", handlers.Make(handlers.HandleHome))
	r.Get("/foo", auth.RequireAuth(handlers.Make(handlers.HandleFoo), authProvider))

	r.Get("/auth/{provider}", handlers.Make(handler.HandleProviderLogin))
	r.Get("/auth/{provider}/callback", handlers.Make(handler.HandlerAuthCallbackFunc))
	r.Get("/auth/logout/{provider}", handlers.Make(nil))
	r.Get("/auth/login", handlers.Make(handler.HandleLogin))

	listenAddr := os.Getenv("LISTEN_ADDR")
	http.ListenAndServe(listenAddr, r)
}
