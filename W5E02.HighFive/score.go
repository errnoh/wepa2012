package main

import (
	"time"
)

var scorecounter Counter

type Score struct {
	Id        int       `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Nickname  string    `json:"nickname"`
	Points    int       `json:"points"`
	Game      *Game     `json:"-"`
}

var scores = make(map[*Game]map[int]*Score)

func addScore(s *Score) {
	var gamescores map[int]*Score
	var exists bool

	if gamescores, exists = scores[s.Game]; !exists {
		gamescores = make(map[int]*Score)
		scores[s.Game] = gamescores
	}

	gamescores[s.Id] = s
}

// Satisfy "sortable" interface
func (s *Score) ID() int {
	return s.Id
}
