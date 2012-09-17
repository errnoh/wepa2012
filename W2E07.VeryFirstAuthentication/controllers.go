package main

import (
	"code.google.com/p/gorilla/sessions"
	"net/http"
	"strings"
)

// FrontController implements http.Handler interface
// so all connections can be directly routed  to it.
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
	println("Added " + path)
	fc.pages[path] = ctrl
}

func (fc *FrontController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var controller Controller
	var ok bool

	if controller, ok = fc.pages[r.URL.Path]; !ok {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	address := controller.processRequest(w, r)
	if strings.HasPrefix(address, "redirect:") {
		address = "/" + strings.SplitAfterN(address, ":", 2)[1]
		http.Redirect(w, r, address, http.StatusFound)
		return

	}

	// html/template oletuksena estää injektiot
	t := getTemplate(address)
	t.Execute(w, r)
}

var store = sessions.NewCookieStore([]byte("baconbaconbacon"))

//----------

type Controller interface {
	processRequest(w http.ResponseWriter, r *http.Request) string
}

type IndexController struct{}

func (c *IndexController) processRequest(w http.ResponseWriter, r *http.Request) string {
	return "login.html"
}

type LoginController struct {
}

func (c *LoginController) processRequest(w http.ResponseWriter, r *http.Request) string {
	r.ParseForm()
	session, _ := store.Get(r, "session")

	user := r.FormValue("username")
	pass := r.FormValue("password")

	if pass != "secret" {
		return "redirect:"
	}

	session.Values["username"] = user
	session.Values["password"] = pass

	session.Save(r, w)

	return "redirect:secret"
}

type LogoutController struct {
}

func (c *LogoutController) processRequest(w http.ResponseWriter, r *http.Request) string {
	session, _ := store.Get(r, "session")

	delete(session.Values, "username")
	delete(session.Values, "password")

	session.Save(r, w)

	return "redirect:"
}

type SecretController struct {
}

func (c *SecretController) processRequest(w http.ResponseWriter, r *http.Request) string {
	session, _ := store.Get(r, "session")

	if pass, ok := session.Values["password"]; !ok || pass != "secret" {
		return "redirect:"
	}

	return "secret.html"
}
