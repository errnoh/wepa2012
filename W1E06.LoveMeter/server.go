// W1E06
//
// Server that counts a number based on submitted form and displays it. address:port/love
package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math"
	"net/http"
)

func init() {
	http.HandleFunc("/love", func(w http.ResponseWriter, r *http.Request) {
		// Read the parameters into r.Form
		r.ParseForm()

		result := &Results{Name1: r.Form.Get("name1"), Name2: r.Form.Get("name2")}

		if result.Name1 != "" && result.Name2 != "" {
			page, err := template.ParseFiles("result.html")
			if err != nil {
				log.Fatalln(err)
			}
			result.Result = match(result.Name1, result.Name2)
			page.Execute(w, result)
			return
		}

		form, err := ioutil.ReadFile("form.html")
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Fprintf(w, string(form))
	})
}

type Results struct {
	Name1  string
	Name2  string
	Result int
}

func match(name1, name2 string) int {
	var uname1, uname2 = []rune(name1), []rune(name2)
	var minlength, result int

	minlength = int(math.Min(float64(len(uname1)), float64(len(uname2))))
	for i := 0; i < minlength; i++ {
		result += int(uname1[i] * uname2[i])
	}

	return (result % 100) + 42
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
