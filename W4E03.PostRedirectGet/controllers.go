package main

import (
	"bufio"
	"code.google.com/p/gorilla/mux"
	"code.google.com/p/gorilla/schema"
	"code.google.com/p/gorilla/sessions"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/order", OrderHandler)
	r.HandleFunc("/order/{id}", ViewHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}

var store = sessions.NewCookieStore([]byte("randomsecrets"))

func Reply(file string, w http.ResponseWriter, r *http.Request, session *sessions.Session) {
	buf := bufio.NewWriter(w)
	t, _ := template.ParseFiles(file)
	t.Execute(buf, session)
	session.Save(r, w)
	buf.Flush()
}

var itemlist = []string{"Bacon", "Ham", "More bacon"}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	PreventCaching(w)

	session, _ := store.Get(r, "session-stuff")
	vars := mux.Vars(r)

	order := orders[vars["id"]]
	session.Values["Name"] = order.Name
	session.Values["Items"] = order.Items
	session.Values["Address"] = order.Address

	Reply("vieworder.jsp", w, r, session)
}

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	PreventCaching(w)

	session, _ := store.Get(r, "session-stuff")
	if r.Method == "GET" {
		session.Values["possibleitems"] = itemlist
		Reply("form.html", w, r, session)
		return
	}
	r.ParseForm()

	// Fill the order struct
	order := NewOrder()
	decode := schema.NewDecoder()
	decode.Decode(order, r.Form)
	order.ParseItems(r)

	errlist := order.Check()
	if errlist != nil {
		// Add flash warnings
		for _, err := range errlist {
			session.AddFlash(err.Error())
		}
		// Remember the previous form data
		session.Values["Name"] = order.Name
		session.Values["Address"] = order.Address
		for _, v := range order.Items {
			session.Values[v] = "checked"
		}
		session.Save(r, w)

		http.Redirect(w, r, "/order", 302)
		return
	}
	AddOrder(order)

	session.AddFlash("Thanks for your order")
	session.Save(r, w)

	http.Redirect(w, r, fmt.Sprintf("/order/%s", order.Id), 302)
}

func PreventCaching(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
}
