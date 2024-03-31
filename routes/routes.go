package routes

import (
	"Documents/Work/go/first/hello-go/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/books", controllers.FindBooks)
	router.POST("/books", controllers.CreateBook)
}
