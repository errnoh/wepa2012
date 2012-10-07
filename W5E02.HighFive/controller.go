package main

import (
	"code.google.com/p/gorilla/mux"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"
)

func main() {
	r := createMux()
	log.Fatal(http.ListenAndServe(":8080", r))
}

func createMux() (r *mux.Router) {
	r = mux.NewRouter()

	r.HandleFunc("/app/games", NewGameHandler).Methods("POST")
	r.HandleFunc("/app/games", ListGamesHandler).Methods("GET")
	r.HandleFunc("/app/games/{name}", ShowGameHandler).Methods("GET")
	r.HandleFunc("/app/games/{name}", DeleteGameHandler).Methods("DELETE")

	r.HandleFunc("/app/games/{name}/scores", CreateScoreHandler).Methods("POST")
	r.HandleFunc("/app/games/{name}/scores", ListScoresHandler).Methods("GET")
	r.HandleFunc("/app/games/{name}/scores/{id:[0-9]+}", ShowScoreHandler).Methods("GET")
	r.HandleFunc("/app/games/{name}/scores/{id:[0-9]+}", DeleteScoreHandler).Methods("DELETE")

	return
}

// TODO: checkien refaktorointi omiksi funktioikseen.

// Games

func NewGameHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// Read body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Couldn't read request body: %s", err.Error())
		http.Error(w, "Couldn't read request", http.StatusBadRequest)
		return
	}

	// Read the input into a Game struct
	g := new(Game)
	if err = json.Unmarshal(body, &g); err != nil {
		http.Error(w, "Input is not valid json", http.StatusBadRequest)
		return
	}

	// Make sure there's a name
	if g.Name == "" {
		http.Error(w, "No name given", http.StatusBadRequest)
		return
	}
	if _, exists := games[g.Name]; exists {
		http.Error(w, "Game with that name already exists", http.StatusBadRequest)
		return
	}
	g.Id = gamecounter.Get()
	games[g.Name] = g

	// Response with the created game
	response, _ := json.Marshal(&g)
	fmt.Fprint(w, string(response))
}

func ListGamesHandler(w http.ResponseWriter, r *http.Request) {
	if len(games) == 0 {
		fmt.Fprint(w, "")
		return
	}

	var temp []sortable
	for _, game := range games {
		temp = append(temp, game)
	}
	sort.Sort(sortableSlice(temp))
	response, _ := json.Marshal(temp)
	fmt.Fprint(w, string(response))
}

func ShowGameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	game, exists := games[vars["name"]]
	if !exists {
		http.Error(w, "Game doesn't exist", http.StatusBadRequest)
		return
	}
	response, _ := json.Marshal(game)
	fmt.Fprint(w, string(response))
}

func DeleteGameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, exists := games[vars["name"]]
	if !exists {
		http.Error(w, "Game doesn't exist", http.StatusBadRequest)
		return
	}
	delete(games, vars["name"])
	fmt.Fprint(w, "")
}

// Scores

func CreateScoreHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	vars := mux.Vars(r)
	game, exists := games[vars["name"]]
	if !exists {
		http.Error(w, "Game doesn't exist", http.StatusBadRequest)
		return
	}

	// Read body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Couldn't read request body: %s", err.Error())
		http.Error(w, "Couldn't read request", http.StatusBadRequest)
		return
	}

	// Read the input into a Score struct
	score := new(Score)
	if err = json.Unmarshal(body, &score); err != nil {
		http.Error(w, "Input is not valid json", http.StatusBadRequest)
		return
	}

	if score.Nickname == "" {
		http.Error(w, "No nickname given", http.StatusBadRequest)
		return
	}
	score.Id = scorecounter.Get()
	score.Timestamp = time.Now()
	score.Game = game
	addScore(score)

	// Response with the created game
	response, _ := json.Marshal(&score)
	fmt.Fprint(w, string(response))

}

func ListScoresHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	vars := mux.Vars(r)
	game, exists := games[vars["name"]]
	if !exists {
		http.Error(w, "Game doesn't exist", http.StatusBadRequest)
		return
	}
	gamescores, exists := scores[game]
	if !exists {
		fmt.Fprint(w, "[]")
		return
	}

	var temp []sortable
	for _, score := range gamescores {
		temp = append(temp, score)
	}
	sort.Sort(sortableSlice(temp))
	response, _ := json.Marshal(temp)
	fmt.Fprint(w, string(response))
}

func ShowScoreHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	game, exists := games[vars["name"]]
	if !exists {
		http.Error(w, "Game doesn't exist", http.StatusBadRequest)
		return
	}
	gamescores, exists := scores[game]
	if !exists {
		http.Error(w, "Game doesn't have any scores", http.StatusBadRequest)
		return
	}
	if vars["id"] == "" {
		http.Error(w, "No score id specified", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Score id should be a number", http.StatusBadRequest)
		return
	}
	score, exists := gamescores[id]
	if !exists {
		http.Error(w, "Game doesn't have score with that id", http.StatusBadRequest)
		return
	}
	response, _ := json.Marshal(score)
	fmt.Fprint(w, string(response))
}

func DeleteScoreHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	game, exists := games[vars["name"]]
	if !exists {
		http.Error(w, "Game doesn't exist", http.StatusBadRequest)
		return
	}
	gamescores, exists := scores[game]
	if !exists {
		http.Error(w, "Game doesn't have any scores", http.StatusBadRequest)
		return
	}
	if vars["id"] == "" {
		http.Error(w, "No score id specified", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Score id should be a number", http.StatusBadRequest)
		return
	}
	_, exists = gamescores[id]
	if !exists {
		http.Error(w, "Game doesn't have score with that id", http.StatusBadRequest)
		return
	}
	delete(gamescores, id)
	fmt.Fprint(w, "")
}
