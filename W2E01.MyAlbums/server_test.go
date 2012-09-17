package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestList(t *testing.T) {
	var res *http.Response
	var page []byte
	var err error

	// Create a test server and check that the page can be found etc
	ts := httptest.NewServer(http.DefaultServeMux)
	defer ts.Close()
	if res, err = http.Get(ts.URL + "/list"); err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Error("Status code:", res.StatusCode)
	}
	if page, err = ioutil.ReadAll(res.Body); err != nil {
		t.Fatal(err)
	}

	site := string(page)

	for _, album := range albums {
		if !strings.Contains(site, album.Name) {
			println(site)
			t.Errorf("Missing %s", album.Name)
		}
		for _, song := range album.Tracks {
			if !strings.Contains(site, song) {
				t.Errorf("Missing %s", album.Name)
			}
		}
	}
}
