package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {

	router := httprouter.New()
	router.HandlerFunc(http.MethodPut, "/next", app.NextTask)
	router.HandlerFunc(http.MethodGet, "/task/:id", app.Task)
	// Return the httprouter instance.
	return app.enableCORS(router)

	// mux := pat.New()
	// mux.Get("/login", http.HandlerFunc(app.Login))
	// mux.Get("/aboutus", http.HandlerFunc(app.AboutUs))
	// mux.Get("/privacypolicy", http.HandlerFunc(app.PrivacyPolicy))
	// mux.Get("/task/:id", http.HandlerFunc(app.Task))
	// //mux.Get("/checkUser", http.HandlerFunc(app.checkUser))
	// mux.Get("/techcontent", http.HandlerFunc(app.TechContent))
	// mux.Put("/next", http.HandlerFunc(app.NextTask))
	// mux.Get("/auth/:provider/callback", http.HandlerFunc(app.auth))
	// mux.Get("/auth/:provider", http.HandlerFunc(gothic.BeginAuthHandler))
	// mux.Get("/", http.HandlerFunc(app.showTemplates))
	// return app.enableCORS(mux)
}
