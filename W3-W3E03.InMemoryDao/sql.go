package main

import (
	"fmt"
	"github.com/errnoh/dependency"
)

func init() {
	dependency.SetConfig("", "text")
	dependency.Add("dao", "mapdao", NewMapDAO)
	dependency.Refresh()
}

func main() {
	app := NewAlbumApp()
	fmt.Println(app)

	app.populateAlbums()
	fmt.Println(app)

	app.deleteAlbumsFromArtist("Pekonimaa")
	fmt.Println(app)
}
