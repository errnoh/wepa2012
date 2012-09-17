package main

import (
	"html/template"
	"log"
	"net/http"
)

func getTemplate(file string) *template.Template {
	page, err := template.ParseFiles(file)
	if err != nil {
		log.Fatalln(err)
	}
	return page
}

func main() {
	controller := NewFrontController()

	controller.AddController("/new-password", &PasswordGeneratorController{})

	log.Fatal(http.ListenAndServe(":8080", controller))
}
