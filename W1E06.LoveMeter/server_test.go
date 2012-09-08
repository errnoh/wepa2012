package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestForm(t *testing.T) {
	var res *http.Response
	var got []byte
	var err error

	ts := httptest.NewServer(http.DefaultServeMux)
	defer ts.Close()
	if res, err = http.Get(ts.URL + "/love?name1=mikke&name2=kasper"); err != nil {
		t.Error(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		t.Error("Status code:", res.StatusCode)
	}
	if got, err = ioutil.ReadAll(res.Body); err != nil {
		t.Error(err)
	}
	if !strings.Contains(string(got), "mikke and kasper match 80%") {
		t.Error("Page should contain text \"mikke and kasper match 80%\" but it doesn't:\n", string(got))
	}
}

func TestMatch(t *testing.T) {
	percent := match("mikke", "kasper")
	if percent != 80 {
		t.Errorf("Mikke and Kasper should be 80%% match, instead they got %d%%", percent)
	}
}
