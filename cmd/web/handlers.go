package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gopheramit/Learning-Python/pkg/models"
	"github.com/julienschmidt/httprouter"
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

func (app *application) GetUserById(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	tid := params.ByName("userid")
	snil := &models.PythonUser{
		UserID: "",
		Email:  "",
		TaskID: 1,
	}

	s, err := app.users.GetID(tid)
	if s != nil {
		app.logger.Println("inserting new data")
		if err != nil {

			w.WriteHeader(http.StatusInternalServerError)
		}

		err := app.writeJSON(w, http.StatusOK, s, nil)
		if err != nil {
			app.serverErrorResponse(w, r, err)
		}
	} else {
		app.writeJSON(w, http.StatusOK, snil, nil)
	}
}

func (app *application) PostUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("ccess-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	t := &models.PythonUser{}
	//w.Header().Set("Access-Control-Allow-Origin", "*")

	err := app.readJSON(w, r, t)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	id, err := app.users.Insert(t.UserID, t.Email, t.TaskID)
	if err != nil {
		fmt.Println("Error occured during insert to database")
	}
	app.logger.Println(id)
	app.writeJSON(w, http.StatusOK, err, nil)

}

func (app *application) auth(res http.ResponseWriter, req *http.Request) {
	log.Println("In AUTH")
	user, err := gothic.CompleteUserAuth(res, req)

	if err != nil {
		log.Println(err)
		return
	}
	s, err := app.users.GetID(user.UserID)
	if s != nil {
		app.logger.Println("inserting new data")
		if err != nil {

			res.WriteHeader(http.StatusInternalServerError)
		}

		return
	} else {
		id, err := app.users.Insert(user.UserID, user.Email, 1)
		fmt.Println(user.UserID)

		if err != nil {
			fmt.Println("Error occured during insert to database")
		}
		//fmt.Println(id)
		app.logger.Println(id)

	}
	err = app.writeJSON(res, http.StatusOK, s, nil)
	if err != nil {
		app.serverErrorResponse(res, req, err)
	}
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//var tid int = 1
	// tid, err := strconv.Atoi(r.URL.Query().Get(":id"))
	params := httprouter.ParamsFromContext(r.Context())
	tid, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || tid < 1 {
		app.notFoundResponse(w, r)
		return
	}

	task, _ := app.users.GetTaskByID(int(tid))

	if task != nil {
		err := app.writeJSON(w, http.StatusOK, task, nil)
		if err != nil {
			app.serverErrorResponse(w, r, err)
		}
	} else {
		app.logger.Println("Task for taskid see below not found")
		app.logger.Println(tid)

	}
}

func (app *application) TechContent(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello from Tech Content"))
}

func (app *application) NextTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("ccess-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	t := &models.PythonUser{}
	//w.Header().Set("Access-Control-Allow-Origin", "*")

	err := app.readJSON(w, r, t)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}
	taskId := t.TaskID
	userId := t.UserID
	tid, _ := app.users.IncrementTaskId(taskId, userId)
	task, _ := app.users.GetTaskByID(tid)

	if task != nil {
		err := app.writeJSON(w, http.StatusOK, task, nil)
		if err != nil {
			app.serverErrorResponse(w, r, err)
		}
	} else {
		app.logger.Println("Task for taskid see below not found")
		app.logger.Println(tid)

	}
}
