package main

import (
	"strconv"
)

var gulls = make(map[string]*Gull)

type Gull struct {
	Id       string `json:"id"`
	Location string `json:"location"`
}

var counter int

func NewGull() *Gull {
	counter++
	return &Gull{Id: strconv.Itoa(counter)}
}

func AddGull(gull *Gull) {
	gulls[gull.Id] = gull
}

func RemoveGull(id string) bool {
	if _, ok := gulls[id]; ok {
		delete(gulls, id)
		return ok
	}
	return false
}

func GetGulls() []*Gull {
	var all = make([]*Gull, 0, 10)
	for _, gull := range gulls {
		all = append(all, gull)
	}
	return all
}
