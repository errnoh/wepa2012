package main

import (
	"bufio"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestForm(t *testing.T) {
	var res *http.Response
	var page []byte
	var err error
	var line []byte

	// Create a test server and check that the page can be found etc
	ts := httptest.NewServer(http.DefaultServeMux)
	defer ts.Close()
	if res, err = http.Get(ts.URL + "/form"); err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Error("Status code:", res.StatusCode)
	}
	if page, err = ioutil.ReadAll(res.Body); err != nil {
		t.Fatal(err)
	}

	// Check that a line matching each of these is found on the page.
	// The fields can be in any order but must be on the same line.
	var contents = [][]string{
		[]string{`name="name"`, `id="name"`, `type="text"`},
		[]string{`name="address"`, `id="address"`, `type="text"`},
		[]string{`id="ticket-green"`, `name="ticket"`, `type="radio"`, `value="green"`},
		[]string{`id="ticket-yellow"`, `name="ticket"`, `type="radio"`, `value="yellow"`},
		[]string{`id="ticket-red"`, `name="ticket"`, `type="radio"`, `value="red"`},
		[]string{`type="submit"`, `value="Tilaa"`},
	}

	reader := bufio.NewReader(strings.NewReader(string(page)))

	var found int
	for {
		line, _, err = reader.ReadLine()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
		}
	checkloop:
		for i, slice := range contents {
			for j, field := range slice {
				if !strings.Contains(string(line), field) {
					continue checkloop
				}
				if j == len(slice)-1 {
					contents[i] = nil
					found++
				}
			}
			if found == len(contents) {
				return
			}
		}
	}
	t.Errorf("Missing following items from form:\n%s", contents)

}
