package main

import (
	"github.com/nightfury1204/demo-macaron/pkg"
	"gopkg.in/macaron.v1"
	"net/http"
)

func main() {
	store := pkg.NewInmemStorage()

	m := macaron.Classic()

	m.Use(macaron.Renderer())

	m.Get("/", func(ctx *macaron.Context) {
		ctx.HTML(http.StatusOK, "home")
	})

	m.Group("/books", func() {
		m.Combo("").Get(store.GetAllBooks). // get all books
			Post(store.CreateBook) // add new book

		m.Combo("/:id").
			Get(store.GetBook). // get the specific book
			Put(store.EditBook). // update book info
			Delete(store.DeleteBook) // delete book
	})

	m.Run()
}
