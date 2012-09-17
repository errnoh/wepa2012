// W201
//
// List Album names and tracks
package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Album struct {
	Artist string
	Name   string
	Tracks []string
}

func GenerateAlbums() (albums []*Album) {
	var album *Album

	album = &Album{
		Artist: "rendezvous park",
		Name:   "closer to being here",
		Tracks: []string{"soothe", "closer to being here", "ascension"},
	}
	albums = append(albums, album)

	album = &Album{
		Artist: "rendezvous park",
		Name:   "the days you didnt notice",
		Tracks: []string{"and grey was the morning...",
			"fell asleep on my lap (eternally)",
			"down - part I",
			"down - part II",
			"anything but hopeless",
			"unnoticed beginnings",
			"...leaves",
		},
	}
	albums = append(albums, album)

	return
}

var albums []*Album

func getTemplate(file string) *template.Template {
	page, err := template.ParseFiles(file)
	if err != nil {
		log.Fatalln(err)
	}
	return page
}

func init() {
	albums = GenerateAlbums()

	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		listpage := getTemplate("list.html")
		listpage.Execute(w, albums)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		form, err := ioutil.ReadFile("index.html")
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Fprintf(w, string(form))
	})
}

func main() {
	fmt.Println(":8080/list and :8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
