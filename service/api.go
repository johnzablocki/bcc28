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

	//define the routes
	router.GET("/", api.BooksHandler.List)   //http://localhost:8888/
	router.POST("/", api.BooksHandler.Post)  //http://localhost:8888/
	router.GET("/:id", api.BooksHandler.Get) //http://localhost:8888/17
	return router.Run(fmt.Sprintf(":%d", config.Port))
}

//NewAPI factory method for API
func NewAPI(handler *BooksHandler) *API {
	return &API{
		BooksHandler: handler,
	}
}
