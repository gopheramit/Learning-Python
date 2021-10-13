package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

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
	//mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/login", http.HandlerFunc(app.Login))
	mux.Get("/aboutus", http.HandlerFunc(app.AboutUs))
	mux.Get("/privacypolicy", http.HandlerFunc(app.PrivacyPolicy))
	mux.Get("/task", http.HandlerFunc(app.Task))
	mux.Get("/techcontent", http.HandlerFunc(app.TechContent))

	mux.Get("/auth/:provider/callback", http.HandlerFunc(app.auth))
	mux.Get("/auth/:provider", http.HandlerFunc(gothic.BeginAuthHandler))
	mux.Get("/", http.HandlerFunc(app.showTemplates))

	return mux
}

func beginAuth(res http.ResponseWriter, req *http.Request) {
	gothic.BeginAuthHandler(res, req)
}

func (app *application) showTemplates(res http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("cmd/web/index.html")
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
	}
	t.Execute(res, false)
}

func (app *application) auth(res http.ResponseWriter, req *http.Request) {
	log.Println("In AUTH")
	user, err := gothic.CompleteUserAuth(res, req)

	if err != nil {
		fmt.Fprintln(res, req)
		fmt.Println("error here")
		return
	}

	s, err := app.users.GetID(user.UserID)

	if s != nil {
		// app.session.Put(r, "authenticatedUserID", s.ID)

		t, err := template.ParseFiles("cmd/web/sucess.html")
		if err != nil {
			fmt.Println(err)
			res.WriteHeader(http.StatusInternalServerError)
		}
		log.Println("Parsed the template")
		t.Execute(res, user)
		// http.Redirect(w, r, fmt.Sprintf("/scrap/%d", s.ID), http.StatusSeeOther)
		return
	} else {
		id, err := app.users.Insert(user.UserID, user.Email, 1)
		if err != nil {
			// if errors.Is(err, models.ErrDuplicateEmail) {
			// 	form.Errors.Add("email", "Address is already in use")
			// 	app.render(w, r, "login.page.tmpl", &templateData{Form: form})
			// } else {
			fmt.Println("Error occured during insert to database")
			// }
			// return
		}
		fmt.Println(id)
		// app.session.Put(r, "authenticatedUserID", id)
		t, err := template.ParseFiles("cmd/web/sucess.html")
		if err != nil {
			fmt.Println(err)
			res.WriteHeader(http.StatusInternalServerError)
		}
		log.Println("Parsed the template")
		t.Execute(res, user)
		// http.Redirect(w, r, fmt.Sprintf("/scrap/%d", id), http.StatusSeeOther)
	}
	// if err != nil {
	// 	fmt.Fprintln(res, err)
	// 	return
	// }
	// log.Println("Completed with AUTH")

}
