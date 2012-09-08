// W1E04
//
// Web server that serves a basic form in address:port/form
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func init() {
	// NOTE: Handlereiksi voi syöttää myös funktioita,
	//       mutta tehtävissä esiintyvien servereiden ollessa näin minimalistisia 
	//       käytetään nyt tätä muotoa erillisen funktion sijaan.
	//
	//  Luetaan sivun pohja tiedostosta ja palautetaan se sitä pyytäneelle.
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
