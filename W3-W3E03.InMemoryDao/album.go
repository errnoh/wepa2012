package main

import (
	"fmt"
)

type Album struct {
	id     string
	artist string
	name   string
	year   int
}

func (a *Album) GetID() string {
	return a.id
}

func (a *Album) SetID(id string) {
	a.id = id
}

func (a *Album) String() string {
	return fmt.Sprintf("[%d] %s - %s", a.year, a.artist, a.name)
}

func ToAlbum(o Object) (a *Album) {
	a, _ = o.(*Album)
	return a
}
