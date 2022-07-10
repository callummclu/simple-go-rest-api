package main

import (
	"github.com/callummclu/simple-go-rest-api/controllers"
	"github.com/callummclu/simple-go-rest-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/users", controllers.FindUsers)
	r.GET("/users/:id", controllers.FindUser)
	r.POST("/users", controllers.CreateUser)
	r.PATCH("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	// Run the server
	r.Run()
}
