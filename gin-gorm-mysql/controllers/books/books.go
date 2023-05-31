package books

import (
	"net/http"

	"gin-gorm-mysql/controllers/books/dtos"
	"gin-gorm-mysql/models"

	"github.com/gin-gonic/gin"
)

func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func FindBookByID(c *gin.Context) {
	var book models.Book
	if err := models.DB.First(&book, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func CreateBook(c *gin.Context) {
	var bookData dtos.CreateBookDTO
	if err := c.ShouldBindJSON(&bookData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newBook := models.Book{
		Title:  bookData.Title,
		Author: bookData.Author,
	}

	models.DB.Create(&newBook)

	c.JSON(http.StatusOK, gin.H{"data": newBook})
}

func UpdateBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input dtos.UpdateBookDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBookByID(c *gin.Context) {
	var book models.Book
	if err := models.DB.First(&book, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}
