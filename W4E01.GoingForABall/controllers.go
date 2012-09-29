package main

import (
	"bufio"
	"code.google.com/p/gorilla/schema"
	"code.google.com/p/gorilla/sessions"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/registration", RegisterHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var store = sessions.NewCookieStore([]byte("randomsecrets"))

// Read the site template, fill it into a buffer, save the session and flush the buffer as a reply to the client.
func Reply(file string, w http.ResponseWriter, r *http.Request, session *sessions.Session) {
	buf := bufio.NewWriter(w)
	t, _ := template.ParseFiles(file)
	t.Execute(buf, session)
	session.Save(r, w)
	buf.Flush()
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	PreventCaching(w)

	session, _ := store.Get(r, "session-stuff")
	if r.Method == "GET" {
		Reply("form.html", w, r, session)
		return
	}
	r.ParseForm()

	// Create new Registration and parse the form into it
	reg := new(Registration)
	decode := schema.NewDecoder()
	decode.Decode(reg, r.Form)

	errlist := reg.Check()
	if errlist != nil {
		for _, err := range errlist {
			session.AddFlash(err.Error())
		}
		session.Save(r, w)

		http.Redirect(w, r, "/registration", 302)
		return
	}
	if ContainsRegistration(reg) {
		session.AddFlash(alreadyThereError.Error())
		session.Save(r, w)

		http.Redirect(w, r, "/registration", 302)
		return
	}
	AddRegistration(reg)

	session.AddFlash("Thanks for registration")
	session.Save(r, w)

	http.Redirect(w, r, "/registration", 302)
}

func PreventCaching(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
}
