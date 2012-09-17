package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestList(t *testing.T) {
	var res *http.Response
	var page []byte
	var err error

	controller := addControllers()

	ts := httptest.NewServer(controller)
	defer ts.Close()

	form := make(url.Values)
	for i := 0; i < 300; i++ {
		form.Set("message", fmt.Sprintf("Spam%d", i))
		res, _ = http.PostForm(
			ts.URL+"/add-message",
			form,
		)
		res.Body.Close()
	}

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

	for i := 0; i < 300; i++ {
		if !strings.Contains(site, fmt.Sprintf("Spam%d", i)) {
			t.Fatal("Didn't contain all the messages")
		}

	}
}
