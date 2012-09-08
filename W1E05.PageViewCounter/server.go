// W1E05
//
// Server that displays a page hit counter in address:port/count
package main

import (
	"html/template"
	"log"
	"net/http"
	"sync"
)

// Counter that can only be increased by one thread at a time.
type counter struct {
	value uint
	mutex *sync.Mutex
}

func NewCounter() (c *counter) {
	c = new(counter)
	c.mutex = new(sync.Mutex)
	return
}

// Increases the counter and returns the new value.
func (c *counter) Increase() (u uint) {
	c.mutex.Lock()
	c.value++
	u = c.value
	c.mutex.Unlock()
	return
}

var hitcounter = NewCounter()

// Reads a page template from a file
func getTemplate(file string) *template.Template {
	page, err := template.ParseFiles(file)
	if err != nil {
		log.Fatalln(err)
	}
	return page
}

var countpage = getTemplate("count.html")

func init() {
	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		count := hitcounter.Increase()
		countpage.Execute(w, count)
	})
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
