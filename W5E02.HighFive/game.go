package main

var gamecounter Counter

type Game struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// Satisfy "sortable" interface
func (g *Game) ID() int {
	return g.Id
}

var games = make(map[string]*Game)
