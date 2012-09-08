package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Testailee vain latautuuko sivu.
func TestCount(t *testing.T) {
	var res *http.Response
	var got []byte
	var err error

	ts := httptest.NewServer(http.DefaultServeMux)
	defer ts.Close()
	for i := 0; i < 10170; i++ {
		res, _ := http.Get(ts.URL + "/count")
		res.Body.Close()
	}
	if res, err = http.Get(ts.URL + "/count"); err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Error("Status code:", res.StatusCode)
	}
	if got, err = ioutil.ReadAll(res.Body); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(got), "Laskuri: 10171") {
		t.Fatalf("Expected count to be 10171. This is what we got:\n%s", string(got))
	}

}
