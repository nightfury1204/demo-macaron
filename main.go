package main

import (
	"github.com/nightfury1204/demo-macaron/pkg"
	"log"
	"gopkg.in/macaron.v1"
	"net/http"
)

func main() {
	// store := pkg.NewInmemStorage()
    if err := pkg.InitDBEngine("root:root@tcp(127.0.0.1:3306)/library?charset=utf8"); err != nil {
		log.Fatal(err)
	}

	m := macaron.Classic()

	m.Use(macaron.Renderer())

	m.Get("/", func(ctx *macaron.Context) {
		ctx.HTML(http.StatusOK, "home")
	})

	m.Group("/books", func() {
		m.Combo("").Get(pkg.GetAllBooks). // get all books
			Post(pkg.CreateBook) // add new book

		m.Combo("/:id").
			Get(pkg.GetBook). // get the specific book
			Put(pkg.EditBook). // update book info
			Delete(pkg.DeleteBook) // delete book
	})

	m.Run()
}
