package routes

import (
	"Documents/Work/go/first/hello-go/controllers"
	"Documents/Work/go/first/hello-go/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	// Other rest routes to demonstrate basic db quering - TASK 04
	router.GET("/books", controllers.FindBooks)
	router.GET("/books/:id", controllers.FindBook)
	router.PATCH("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	// GET endpoint to demonstrate custom bindings - TASK 03
	router.GET("/bookable", controllers.GetBookable)

	// POST endpoint with rate limit - TASK 01
	router.POST("/books", middleware.EnsureRateLimit, controllers.CreateBook)
}
