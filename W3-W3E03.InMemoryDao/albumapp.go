package main

import (
	"fmt"
	"github.com/errnoh/dependency"
)

type AlbumApp struct {
	dao DAO
}

func NewAlbumApp() *AlbumApp {
	var err error
	var emptyinterface interface{}

	app := new(AlbumApp)
	if emptyinterface, err = dependency.Get("dao"); err != nil {
		fmt.Println(err.Error())
	}
	app.dao, _ = emptyinterface.(DAO)
	return app
}

func (app *AlbumApp) populateAlbums() {
	albums := []*Album{
		&Album{"", "Pekonimaa", "Polku Itämereen", 1923},
		&Album{"", "Pekonimaa", "Merenpohjaan ja takaisin", 1925},
		&Album{"", "Kaali", "Hapan vai ei?", 1973},
		&Album{"", "Nappikauppiaat", "Kadulla", 1968},
	}
	for _, a := range albums {
		app.dao.create(a)
	}
}

func (app *AlbumApp) deleteAlbumsFromArtist(artist string) {
	for _, a := range app.dao.list() {
		if ToAlbum(a).artist == artist {
			app.dao.remove(a.GetID())
		}
	}
}

func (app *AlbumApp) String() (s string) {
	for _, a := range app.dao.list() {
		s += fmt.Sprintln(ToAlbum(a))
	}
	return s
}
