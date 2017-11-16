package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type (
	//API struct
	API struct {
		BooksHandler *BooksHandler
	}
)

//Init kick off api
func (api *API) Init(config *Config) error {
	router := gin.Default()
	router.GET("/", api.BooksHandler.List)
	router.POST("/", api.BooksHandler.Post)
	router.GET("/:id", api.BooksHandler.Get)
	return router.Run(fmt.Sprintf(":%d", config.Port))
}

//NewAPI factory method for API
func NewAPI(handler *BooksHandler) *API {
	return &API{
		BooksHandler: handler,
	}
}
