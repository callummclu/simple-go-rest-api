package main

import (
	"simple-go-rest-api/models"
	"simple-go-rest-api/routes"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func main() {

	var itemList = []models.Item{
		{Id: 1, Name: "doc 1"},
		{Id: 2, Name: "doc 2"},
	}

	routes.CreateUrlMappings(itemList)
	routes.Router.Run(":8000")
}
