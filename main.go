package main

/*
	This project is a sample app to accompany the presentation
	Going, Going, Golang...  at Boston Code Camp 28.  It is meant
	as an introduction to Go, with some basic patterns.  Extensive
	comments are provided for clarification of basic Go concepts.
*/

import (
	"log"

	"github.com/jzablocki/bcc28/data"
	"github.com/jzablocki/bcc28/service"
)

//Main entry point into the program
func main() {

	//Create a new BookStore via its factory method
	bookStore, err := data.NewBookStore()

	if err != nil {
		//crash the program if a bookstore can't be created
		log.Panic(err)
	}
	bookStore.Init(&data.Config{DBName: "books.db"})

	//Create a new API and pass it its BookStore dependency
	api := service.NewAPI(&service.BooksHandler{
		BookStore: bookStore,
	})

	//Init will configure the service to run and listen until it crashs
	//or is shut down with force
	log.Panic(api.Init(&service.Config{Port: 8888}))
}
