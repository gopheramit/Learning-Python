package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/gorilla/sessions"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func (app *application) routes() http.Handler {
	key1 := "DoN4QZCXaa3TJfr4BJZMQZNo"
	maxAge := 86400 * 30 // 30 days
	isProd := false      // Set to true when serving over https
	store := sessions.NewCookieStore([]byte(key1))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd
	gothic.Store = store
	goth.UseProviders(
		google.New("263741611747-2bgmmh2vnbjvt02c3m8s30ujbb76obgf.apps.googleusercontent.com", "DoN4QZCXaa3TJfr4BJZMQZNo", "http://localhost:4000/auth/google/callback", "email", "profile"),
	)
	mux := pat.New()
	mux.Get("/login", http.HandlerFunc(app.Login))
	mux.Get("/aboutus", http.HandlerFunc(app.AboutUs))
	mux.Get("/privacypolicy", http.HandlerFunc(app.PrivacyPolicy))
	mux.Get("/task/:id", http.HandlerFunc(app.Task))
	//mux.Get("/", http.HandlerFunc(app.Task))
	mux.Get("/techcontent", http.HandlerFunc(app.TechContent))
	mux.Put("/next", http.HandlerFunc(app.NextTask))
	mux.Get("/auth/:provider/callback", http.HandlerFunc(app.auth))
	mux.Get("/auth/:provider", http.HandlerFunc(gothic.BeginAuthHandler))
	mux.Get("/", http.HandlerFunc(app.showTemplates))
	return app.enableCORS(mux)
}
