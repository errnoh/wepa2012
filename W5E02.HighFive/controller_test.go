package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func clear() {
	games = make(map[string]*Game)
	scores = make(map[*Game]map[int]*Score)
	gamecounter = Counter(0)
	scorecounter = Counter(0)
}

func postGame(s *httptest.Server, i int) {
	body, _ := json.Marshal(map[string]interface{}{
		"name": fmt.Sprintf("testbacon-%d", i),
	})
	http.Post(s.URL+"/app/games", "application/json", strings.NewReader(string(body)))
}

func postScore(s *httptest.Server, game int, name string) {
	body, _ := json.Marshal(map[string]interface{}{
		"nickname": name,
		"points":   game,
	})
	http.Post(fmt.Sprintf("%s/app/games/testbacon-%d/scores", s.URL, game), "application/json", strings.NewReader(string(body)))
}

func TestCreateGame(t *testing.T) {
	clear()
	r := createMux()
	s := httptest.NewServer(r)

	body, err := json.Marshal(map[string]interface{}{
		"name": "testbacon",
	})
	if err != nil {
		t.Errorf("Couldn't marshal input: %s", err.Error())
	}

	_, err = http.Post(s.URL+"/app/games", "application/json", strings.NewReader(string(body)))
	if err != nil {
		t.Fatalf("Failed to POST: %s", err.Error())
	}

	if _, exists := games["testbacon"]; !exists {
		t.Fatalf("Failed to add a game: %+v", games)
	}

}

func ExampleListGames() {
	clear()
	r := createMux()
	s := httptest.NewServer(r)

	for i := 0; i < 10; i++ {
		postGame(s, i)
	}
	resp, err := http.Get(s.URL + "/app/games")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Print(string(body))
	// Output:
	// [{"id":1,"name":"testbacon-0"},{"id":2,"name":"testbacon-1"},{"id":3,"name":"testbacon-2"},{"id":4,"name":"testbacon-3"},{"id":5,"name":"testbacon-4"},{"id":6,"name":"testbacon-5"},{"id":7,"name":"testbacon-6"},{"id":8,"name":"testbacon-7"},{"id":9,"name":"testbacon-8"},{"id":10,"name":"testbacon-9"}]
}

func ExampleShowGame() {
	clear()
	r := createMux()
	s := httptest.NewServer(r)

	postGame(s, 0)

	resp, err := http.Get(s.URL + "/app/games/testbacon-0")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Print(string(body))
	// Output:
	// {"id":1,"name":"testbacon-0"}
}

func ExampleDeleteGame() {
	clear()
	r := createMux()
	s := httptest.NewServer(r)

	postGame(s, 0)

	fmt.Println("start:", len(games))
	req, err := http.NewRequest("DELETE", s.URL+"/app/games/testbacon-0", nil)
	if err != nil {
		fmt.Println(err)
	}
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("end:", len(games))
	// Output:
	// start: 1
	// end: 0
}

func ExampleAddScore() {
	clear()
	r := createMux()
	s := httptest.NewServer(r)

	postGame(s, 0)
	postScore(s, 0, "baconman")
	for _, games := range scores {
		for _, score := range games {
			fmt.Println(score.Id, score.Nickname)
		}
	}
	// Output: 1 baconman
}

func TestListScores(t *testing.T) {
	clear()
	r := createMux()
	s := httptest.NewServer(r)

	postGame(s, 0)
	for i := 0; i < 10; i++ {
		postScore(s, 0, fmt.Sprintf("baconman%d", i))
	}
	resp, err := http.Get(s.URL + "/app/games/testbacon-0/scores")
	if err != nil {
		t.Fatalf("Failed to get scores: %s", err.Error())
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	for i := 0; i < 10; i++ {
		if !strings.Contains(string(body), fmt.Sprintf("baconman%d", i)) {
			t.Fatalf("Listing didn't contain baconman%d:\n%s", i, string(body))
		}
	}
}

func TestShowScore(t *testing.T) {
	clear()
	r := createMux()
	s := httptest.NewServer(r)

	postGame(s, 0)
	for i := 0; i < 10; i++ {
		postScore(s, 0, fmt.Sprintf("baconman%d", i))
	}
	resp, err := http.Get(s.URL + "/app/games/testbacon-0/scores/2")
	if err != nil {
		t.Fatalf("Failed to get scores: %s", err.Error())
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if !strings.Contains(string(body), "baconman1") {
		t.Fatalf("Score didn't contain baconman2:\n%s", string(body))
	}
}

func ExampleDeleteScore() {
	clear()
	r := createMux()
	s := httptest.NewServer(r)

	postGame(s, 0)
	postScore(s, 0, "baconman")

	fmt.Println("start:", len(scores[games["testbacon-0"]]))
	req, err := http.NewRequest("DELETE", s.URL+"/app/games/testbacon-0/scores/1", nil)
	if err != nil {
		fmt.Println(err)
	}
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("end:", len(scores[games["testbacon-0"]]))
	// Output:
	// start: 1
	// end: 0
}
