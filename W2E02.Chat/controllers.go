package main

import (
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

	PreventCaching(w)

	if controller, ok = fc.pages[r.URL.Path]; !ok {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	address := controller.processRequest(r)
	if strings.HasPrefix(address, "redirect:") {
		address = "/" + strings.SplitAfterN(address, ":", 2)[1]
		http.Redirect(w, r, address, http.StatusFound)
		return

	}

	// html/template oletuksena estää injektiot
	t := getTemplate(address)
	t.Execute(w, r)
}

func PreventCaching(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
}

//----------

type Controller interface {
	processRequest(r *http.Request) string
}

type IndexController struct{}

func (c *IndexController) processRequest(r *http.Request) string {
	return "redirect:list"
}

type ListMessagesController struct {
	messageService *MessageService
}

func (c *ListMessagesController) processRequest(r *http.Request) string {
	r.ParseForm()

	for _, message := range c.messageService.Messages() {
		r.Form.Add(http.CanonicalHeaderKey("Message"), message)
	}
	return "list.html"

}

type AddMessageController struct {
	messageService *MessageService
}

func (c *AddMessageController) processRequest(r *http.Request) string {
	r.ParseForm()
	message := r.FormValue("message")
	if message != "" {
		c.messageService.Add(message)
	}
	return "redirect:list"
}
