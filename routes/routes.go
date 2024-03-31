package routes

import (
	"Documents/Work/go/first/hello-go/controllers"
	"Documents/Work/go/first/hello-go/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/books", controllers.FindBooks)
	router.GET("/books/:id", controllers.FindBook)
	router.PATCH("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	// Rate limit the POST endpoint
	router.POST("/books", middleware.EnsureRateLimit, controllers.CreateBook)
}
