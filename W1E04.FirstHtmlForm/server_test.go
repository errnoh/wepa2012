package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Testailee vain latautuuko sivu.
func TestForm(t *testing.T) {
	var res *http.Response
	var got []byte
	var err error

	ts := httptest.NewServer(http.DefaultServeMux)
	defer ts.Close()
	if res, err = http.Get(ts.URL + "/form"); err != nil {
		t.Fatal(err)
	}
	if got, err = ioutil.ReadAll(res.Body); err != nil {
		t.Fatal(err)
	}

	println(string(got))
}
