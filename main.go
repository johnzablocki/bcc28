package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jzablocki/bcc28/data"
)

func main() {
	bookStore, err := data.NewBookStore()
	if err != nil {
		panic(err)
	}
	err = bookStore.Init(&data.Config{DBName: "books.db"})
	if err != nil {
		panic(err)
	}

	err = bookStore.Save(&data.Book{
		Model: data.Model{ID: strconv.Itoa(time.Now().Second())},
		ISBN:  "123456789",
		Title: "Learning Go",
	})
	if err != nil {
		panic(err)
	}

	// var savedBook data.Book
	// err = bookStore.FindByID("test", &savedBook)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(savedBook)

	var books []data.Book
	bookStore.FindAll(&books)
	fmt.Println(books)

	err = bookStore.Shutdown()
	if err != nil {
		panic(err)
	}
}
