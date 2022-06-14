package main

import (
	"pratice-sesi8/controllers"
	"pratice-sesi8/database"

	"github.com/gin-gonic/gin"
)

func init() {
	database.ConnectDB()
}

func main() {
	r := gin.Default()
	orderControllers := controllers.NewOrderController(database.GetDB())
	r.GET("/", func(c *gin.Context) {
		c.String(200, "API Running.")
	})
	r.POST("/orders", orderControllers.CreateOrder)
	r.GET("/orders", orderControllers.GetAllOrder)
	r.PUT("/orders/:id", orderControllers.UpdateOrder)
	r.DELETE("/orders/:id", orderControllers.DeleteOrder)
	r.Run(":1234")
}
