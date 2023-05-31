package idcard

import (
	"gin-gorm-mysql/controllers/idcard/dtos"
	"gin-gorm-mysql/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateIDCard(c *gin.Context) {

	var idCard dtos.CreateIDCardDTO

	if err := c.ShouldBindJSON(&idCard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.First(&models.Person{}, idCard.PersonID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Person not found!"})
		return
	}

	newIDCard := models.IDCard{
		IDNumber: idCard.IDNUmber,
		Person:   models.Person{ID: idCard.PersonID},
	}

	models.DB.Create(&newIDCard)

	c.JSON(http.StatusOK, gin.H{"data": newIDCard})
}

func FindIDCardByPersonID(c *gin.Context) {
	var idCard models.IDCard

	if err := models.DB.Where("person_id = ?", c.Param("id")).First(&idCard).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": idCard})
}
