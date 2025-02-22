package urlsmappings

import (
	"github.com/gauravlad21/book-management-system/controller"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	GET    = "GET"
	POST   = "POST"
	PATCH  = "PATCH"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type urlMap struct {
	Url     string
	Method  string
	Handler func(*gin.Context)
}

var urlsMappings []*urlMap

func GetUrlMaps() []*urlMap {
	return urlsMappings
}

func init() {
	urlsMappings = []*urlMap{
		// testing endpoints
		{Url: "/hello", Method: GET, Handler: controller.Hello},
		{Url: "/swagger/*any", Method: GET, Handler: ginSwagger.WrapHandler(swaggerFiles.Handler)},

		// start from here
		{Url: "/books", Method: POST, Handler: controller.CreateBook},
		{Url: "/books/:id", Method: GET, Handler: controller.ReadBook},
		{Url: "/books", Method: GET, Handler: controller.ReadAllBooks},
		{Url: "/books/:id", Method: PUT, Handler: controller.UpdateBook},
		{Url: "/books/:id", Method: DELETE, Handler: controller.DeleteBook},
	}
}
