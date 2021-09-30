package main

import "net/http"

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Learning Python"))
}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello from Login"))
}

func (app *application) AboutUs(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello from About us"))
}

func (app *application) PrivacyPolicy(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello from Policy"))
}

func (app *application) Task(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello from Task"))
}

func (app *application) TechContent(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello from Tech Content"))
}
