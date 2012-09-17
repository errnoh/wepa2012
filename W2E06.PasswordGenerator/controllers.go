package main

import (
	"code.google.com/p/go-uuid/uuid"
	"net/http"
)

type FrontController struct {
	pages map[string]Controller
}

func NewFrontController() (controller *FrontController) {
	controller = &FrontController{
		pages: make(map[string]Controller),
	}
	return
}

func (fc *FrontController) AddController(path string, ctrl Controller) {
	fc.pages[path] = ctrl
}

func (fc *FrontController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var controller Controller
	var ok bool

	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	if controller, ok = fc.pages[r.URL.Path]; !ok {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	address := controller.processRequest(r)
	t := getTemplate(address)
	t.Execute(w, r)
}

type Controller interface {
	processRequest(r *http.Request) string
}

type PasswordGeneratorController struct{}

func (c *PasswordGeneratorController) processRequest(r *http.Request) string {
	r.ParseForm()

	r.Form.Add("password", uuid.NewRandom().String())
	return "password.html"

}
