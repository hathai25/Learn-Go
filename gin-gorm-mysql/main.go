package main

import (
	"gin-gorm-mysql/controllers/books"
	"gin-gorm-mysql/controllers/idcard"
	"gin-gorm-mysql/controllers/person"
	"gin-gorm-mysql/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	client := r.Group("api")
	{
		bookApi := client.Group("book")
		{
			bookApi.GET("/all", books.FindBooks)
			bookApi.GET("/:id", books.FindBookByID)
			bookApi.POST("/create", books.CreateBook)
			bookApi.PUT("/update/:id", books.UpdateBook)
			bookApi.DELETE("/delete/:id", books.DeleteBookByID)
		}
		personApi := client.Group("person")
		{
			personApi.GET("/all", person.FindPersons)
			personApi.GET("/:id", person.FindPersonByID)
			personApi.POST("/create", person.CreatePerson)
			personApi.PUT("/update/:id", person.UpdatePerson)
			personApi.DELETE("/delete/:id", person.DeletePersonByID)
		}
		idcardApi := client.Group("idcard")
		{
			idcardApi.POST("/create", idcard.CreateIDCard)
			idcardApi.GET("/find/person/:id", idcard.FindIDCardByPersonID)
		}
	}

	models.ConnectDataBase()

	//run port 3001
	r.Run(":3001")

}
