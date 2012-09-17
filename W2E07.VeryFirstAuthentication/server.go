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
	controller := addControllers()

	log.Fatal(http.ListenAndServe(":8080", controller))
}

func addControllers() *FrontController {
	controller := NewFrontController()

	controller.AddController("/", &IndexController{})
	controller.AddController("/login", &LoginController{})
	controller.AddController("/logout", &LogoutController{})
	controller.AddController("/secret", &SecretController{})

	return controller
}
