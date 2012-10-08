package main

import (
	"code.google.com/p/gorilla/mux"
	"code.google.com/p/gorilla/schema"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func main() {
	r := createMux()
	log.Fatal(http.ListenAndServe(":8080", r))
}

func createMux() (r *mux.Router) {
	r = mux.NewRouter()

	r.HandleFunc("/app/observation", ListObservationHandler).Methods("GET")
	r.HandleFunc("/app/observation", AddObservationHandler).Methods("POST")
	r.HandleFunc("/app/observationpoint", ListPointHandler).Methods("GET")
	r.HandleFunc("/app/observationpoint", AddPointHandler).Methods("POST")
	return r
}

func PreventCaching(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
}

func ListObservationHandler(w http.ResponseWriter, r *http.Request) {
	PreventCaching(w)

	// Get a custom struct that holds everything we need inside the template
	LOContainer := newLOHelper(r)
	t, err := template.ParseFiles("observations.jsp")
	if err != nil {
		log.Println(err)
	}
	// Execute the template using that struct
	t.Execute(w, LOContainer)
}

func AddObservationHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	o := new(Observation)
	decode := schema.NewDecoder()
	// Fill the struct with the form data
	decode.Decode(o, r.Form)

	id, err := strconv.Atoi(r.FormValue("observationPointId"))
	if err == nil {
		o.Point = observationpoints[id]
	}

	// check if valid
	if o.Point == nil || (r.FormValue("Celsius") != "0" && o.Celsius == 0) {
		http.Redirect(w, r, "/app/observation", http.StatusFound)
		return
	}

	o.Timestamp = time.Now()
	o.Id = observationcounter.Get()

	observations = append(observations, o)
	http.Redirect(w, r, "/app/observation", 302)
}

func ListPointHandler(w http.ResponseWriter, r *http.Request) {
	LPContainer := newPointHelper()
	PreventCaching(w)
	t, err := template.ParseFiles("points.jsp")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, LPContainer)
}

func AddPointHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	o := new(ObservationPoint)
	decode := schema.NewDecoder()
	decode.Decode(o, r.Form)

	if !o.IsValid(r) {
		http.Redirect(w, r, "/app/observationpoint", http.StatusFound)
		return
	}

	o.Id = pointcounter.Get()
	observationpoints = append(observationpoints, o)
	http.Redirect(w, r, "/app/observationpoint", 303)
}
