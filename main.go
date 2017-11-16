package main

import (
	"github.com/jzablocki/bcc28/data"
	"github.com/jzablocki/bcc28/service"
)

func main() {
	bookStore, err := data.NewBookStore()
	if err != nil {
		panic(err)
	}
	bookStore.Init(&data.Config{DBName: "books.db"})

	api := service.NewAPI(&service.BooksHandler{
		BookStore: bookStore,
	})
	panic(api.Init(&service.Config{Port: 8888}))
}
