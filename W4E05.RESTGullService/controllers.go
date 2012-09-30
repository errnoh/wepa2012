package main

import (
	"code.google.com/p/gorilla/mux"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := createmux()
	log.Fatal(http.ListenAndServe(":8080", r))
}

func createmux() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/gull", ListHandler).Methods("GET")
	r.HandleFunc("/gull/{id}", ViewHandler).Methods("GET")
	r.HandleFunc("/gull/{id}", DeleteHandler).Methods("DELETE")
	r.HandleFunc("/gull", AddHandler).Methods("POST")
	return r
}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	enc.Encode(GetGulls())
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if !RemoveGull(vars["id"]) {
		http.Error(w, "Invalid sighting id.", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, "/", 302)
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gull, exists := gulls[vars["id"]]
	if !exists {
		http.Error(w, "Invalid sighting id.", http.StatusNotFound)
		return
	}
	enc := json.NewEncoder(w)
	enc.Encode(gull)
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	gull := NewGull()
	dec.Decode(&gull)
	AddGull(gull)

	http.Redirect(w, r, fmt.Sprintf("/%s", gull.Id), 302)
}
