package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGulls(t *testing.T) {
	r := createmux()
	ts := httptest.NewServer(r)

	// Post Gull
	body, _ := json.Marshal(map[string]interface{}{"location": "Turku"})
	resp, err := http.Post(ts.URL+"/gull", "application/json", strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()

	// List Gulls
	resp, err = http.Get(ts.URL + "/gull")
	if err != nil {
		t.Fatal(err)
	}
	body, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if string(body) != "[{\"id\":\"1\",\"location\":\"Turku\"}]\n" {
		t.Errorf("Body didn't contain correct location: %s", string(body), body)
	}
}
