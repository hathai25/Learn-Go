package main

import (
	"gin-mysql/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	client := r.Group("api")
	{
		client.GET("/story/:id", controllers.Read)
		client.POST("/story", controllers.Create)
		client.PUT("/story/:id", controllers.Update)
		client.DELETE("/story/:id", controllers.Delete)
	}

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
