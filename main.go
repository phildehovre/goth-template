package main

import (
	"log"
	"net/http"
	"os"

	"goth-template/config"
	"goth-template/handlers"
	"goth-template/services/auth"
	"goth-template/store"

	"github.com/go-chi/chi"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file", err)
	}

	cfg := mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	r := chi.NewMux()

	db, err := store.NewMySQLStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}

	s := store.NewStore(db)

	authProvider := auth.NewAuthService()
	handler := handlers.New(s, authProvider)
	r.Handle("/*", public())

	r.Get("/", handlers.Make(handlers.HandleHome))

	r.Get("/auth/{provider}", handlers.Make(handler.HandleProviderLogin))
	r.Get("/auth/{provider}/callback", handlers.Make(nil))
	r.Get("/auth/logout/{provider}", handlers.Make(nil))
	r.Get("/auth/login", handlers.Make(handler.HandleLogin))

	listenAddr := os.Getenv("LISTEN_ADDR")
	http.ListenAndServe(listenAddr, r)
}
