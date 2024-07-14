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
	// Serve static files from the "public" directory
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	r.Handle("/*", public())
	r.Get("/foo", handlers.Make(handlers.HandleFoo))
	listenAddr := os.Getenv("LISTEN_ADDR")
	http.ListenAndServe(listenAddr, r)
}
