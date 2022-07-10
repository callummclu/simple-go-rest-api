package routes

import (
	"fmt"
	"simple-go-rest-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func CreateUrlMappings(itemList []models.Item) {
	Router = gin.Default()

	api := Router.Group("/api")
	{
		api.GET("getAll", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": itemList,
			})
		})

		api.GET("getOne/:id", func(c *gin.Context) {

			id := c.Params.ByName("id")

			for _, item := range itemList {
				fmt.Println(strconv.Itoa(item.Id) == id)
				if strconv.Itoa(item.Id) == id {
					c.JSON(200, gin.H{
						"message": item,
					})
					return
				}
			}
			c.JSON(404, gin.H{
				"message": "not found",
			})
		})

		api.POST("createOne", func(c *gin.Context) {
			var item models.Item
			c.BindJSON(&item)
			itemList = append(itemList, item)
			c.JSON(200, gin.H{
				"message": itemList,
			})
		})

		api.PATCH("editOne/:id", func(c *gin.Context) {
			id := c.Params.ByName("id")

			var item models.Item
			c.BindJSON(&item)

			for i, items := range itemList {
				if strconv.Itoa(items.Id) == id {

					itemList = append(itemList[:i], itemList[i+1:]...)

					itemList = append(itemList, item)
					c.JSON(200, gin.H{
						"message": itemList,
					})
					return
				}
			}

		})

		api.DELETE("deleteOne/:id", func(c *gin.Context) {

			id := c.Params.ByName("id")

			for i, item := range itemList {
				if strconv.Itoa(item.Id) == id {

					itemList = append(itemList[:i], itemList[i+1:]...)

					c.JSON(200, gin.H{
						"message": itemList,
					})
					return
				}
			}
		})

	}
}
