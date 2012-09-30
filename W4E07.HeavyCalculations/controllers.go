package main

import (
	"code.google.com/p/gorilla/mux"
	"code.google.com/p/gorilla/schema"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", TaskHandler)
	r.HandleFunc("/{id}", ViewHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func Reply(file string, w http.ResponseWriter, r *http.Request, data interface{}) {
	t, _ := template.ParseFiles(file)
	t.Execute(w, data)
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	PreventCaching(w)

	vars := mux.Vars(r)
	task, exists := tasks[vars["id"]]
	if !exists {
		http.Error(w, "Invalid task id.", http.StatusNotFound)
		return
	}
	t, _ := template.ParseFiles("viewtask.jsp")
	t.Execute(w, task)
}

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	PreventCaching(w)

	if r.Method == "GET" {
		t, _ := template.ParseFiles("form.jsp")
		t.Execute(w, nil)
		return
	}
	r.ParseForm()

	task := NewTask()
	decode := schema.NewDecoder()
	decode.Decode(task, r.Form)

	// Käytännössä W04E07 tehtävä on vain siinä onko edessä sana 'go' vai ei :)
	go sendForProcessing(task)

	http.Redirect(w, r, fmt.Sprintf("/%s", task.Id), 302)
}

func PreventCaching(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
}
