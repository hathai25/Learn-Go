package person

import (
	"gin-gorm-mysql/controllers/person/dtos"
	"gin-gorm-mysql/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindPersons(c *gin.Context) {
	var persons []models.Person
	models.DB.Find(&persons)

	c.JSON(http.StatusOK, gin.H{"data": persons})
}

func FindPersonByID(c *gin.Context) {
	var person models.Person

	if err := models.DB.First(&person, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": person})

}

func CreatePerson(c *gin.Context) {
	var person dtos.CreatePersonDTO

	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newPerson := models.Person{
		Name: person.Name,
		Age:  person.Age,
	}

	models.DB.Create(&newPerson)

	c.JSON(http.StatusOK, gin.H{"data": person})
}

func UpdatePerson(c *gin.Context) {
	var person models.Person
	if err := models.DB.First(&person, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var updatePerson dtos.UpdatePersonDTO
	if err := c.ShouldBindJSON(&updatePerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&person).Updates(updatePerson)
	c.JSON(http.StatusOK, gin.H{"data": person})
}

func DeletePersonByID(c *gin.Context) {
	var person models.Person
	if err := models.DB.First(&person, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&person)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
