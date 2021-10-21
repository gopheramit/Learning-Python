package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/markbates/goth/gothic"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/" {
	// 	http.NotFound(w, r)
	// 	return
	// }
	// w.Write([]byte("Hello from Learning Python"))
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
		log.Println(err)
		return
	}
	s, err := app.users.GetID(user.UserID)

	err = app.writeJSON(res, http.StatusOK, s, nil)
	if err != nil {
		app.serverErrorResponse(res, req, err)
	}

	// if s != nil {
	// 	t, err := template.ParseFiles("cmd/web/sucess.html")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		res.WriteHeader(http.StatusInternalServerError)
	// 	}
	// 	log.Println("Parsed the template")
	// 	t.Execute(res, user)
	// 	return
	// } else {
	// 	id, err := app.users.Insert(user.UserID, user.Email, 1)
	// 	fmt.Println(user.UserID)

	// 	if err != nil {
	// 		fmt.Println("Error occured during insert to database")
	// 	}
	// 	fmt.Println(id)
	// 	t, err := template.ParseFiles("cmd/web/sucess.html")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		res.WriteHeader(http.StatusInternalServerError)
	// 	}
	// 	log.Println("Parsed the template")
	// 	t.Execute(res, user)
	// }
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

	var tid int = 1

	task, _ := app.users.GetTaskByID(tid)

	if task != nil {
		t, err := template.ParseFiles("cmd/web/task.page.tmpl")
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		log.Println("Parsed the template task values are " + task.TaskName)
		t.Execute(w, task)
		return
	} else {
		w.Write([]byte("No Task Found"))
	}
}

func (app *application) TechContent(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello from Tech Content"))
}

func (app *application) NextTask(w http.ResponseWriter, r *http.Request) {
	var taskId int = 1

	tid, _ := app.users.IncrementTaskId(taskId, "104831921254349297331")
	task, _ := app.users.GetTaskByID(tid)

	if task != nil {
		t, err := template.ParseFiles("cmd/web/task.page.tmpl")
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		log.Println("Parsed the template task values are " + task.TaskName)
		t.Execute(w, task)
		return
	} else {
		w.Write([]byte("No Task Found"))
	}
}
