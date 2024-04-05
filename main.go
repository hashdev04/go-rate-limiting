package main

import (
	"Documents/Work/go/first/hello-go/controllers/validations"
	"Documents/Work/go/first/hello-go/models"
	"Documents/Work/go/first/hello-go/routes"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	routes.Routes(router)
	return router
}

func main() {
	router := setupRouter()
	models.ConnectDatabase()
	validations.Register()
	router.Run(":8080")
}
