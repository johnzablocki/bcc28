package service

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jzablocki/bcc28/data"
)

type (
	//BooksHandler handles book requests
	BooksHandler struct {
		BookStore *data.BookStore
	}
)

//List show all books
func (h *BooksHandler) List(c *gin.Context) {
	var books []data.Book
	h.BookStore.FindAll(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}

//Get find a book by ID
func (h *BooksHandler) Get(c *gin.Context) {
	id := c.Param("id")
	var book data.Book
	h.BookStore.FindByID(id, &book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

//Post create a book
func (h *BooksHandler) Post(c *gin.Context) {
	var book data.Book
	c.BindJSON(&book)
	book.ID = strconv.Itoa(time.Now().Second())
	h.BookStore.Save(&book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}
