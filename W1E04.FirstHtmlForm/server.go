package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/display", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.New("displaystuff").ParseFiles("index.html")
		if err != nil {
			log.Println(err)
			return
		}
		tmpl = tmpl

		fmt.Fprint(w, "InsertStuffHere")
	})

	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		form, err := ioutil.ReadFile("form.html")
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Fprintf(w, string(form))
	})
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
