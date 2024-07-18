package auth

import "github.com/gorilla/sessions"

type SessionsOptions struct {
	CookiesKey string
	MaxAge     int
	HttpOnly   bool
	Secure     bool
}

func NewCookieStore(opts SessionsOptions) *sessions.CookieStore {
	store := sessions.NewCookieStore([]byte(""))
	store.MaxAge(opts.MaxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = opts.HttpOnly
	store.Options.Secure = opts.Secure

	return store
}
