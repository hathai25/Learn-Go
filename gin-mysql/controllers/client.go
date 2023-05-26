package controllers

import (
	"gin-mysql/database"

	"github.com/gin-gonic/gin"
)

type Post struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"body"`
}

func Read(c *gin.Context) {
	db := database.DBConn()
	rows, err := db.Query("SELECT * FROM post WHERE id = " + c.Param("id"))

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error",
		})
	}

	post := Post{}

	for rows.Next() {
		var id int
		var title, body string

		err = rows.Scan(&id, &title, &body)
		if err != nil {
			panic(err.Error())
		}

		post.Id = id
		post.Title = title
		post.Content = body
	}

	c.JSON(200, post)
	defer db.Close()
}

func Create(c *gin.Context) {
	db := database.DBConn()

	type CreatePost struct {
		Title string `form:"title" json:"title" binding:"required"`
		Body  string `form:"body" json:"body" binding:"required"`
	}

	var json CreatePost

	if err := c.ShouldBindJSON(&json); err == nil {
		insPost, err := db.Prepare("INSERT INTO post(title, content) VALUES(?,?)")
		if err != nil {
			c.JSON(500, gin.H{
				"messages": err,
			})
		}

		insPost.Exec(json.Title, json.Body)
		c.JSON(200, gin.H{
			"messages": "inserted",
		})

	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	defer db.Close()
}

func Update(c *gin.Context) {
	db := database.DBConn()
	defer db.Close()

	type UpdatePost struct {
		Title string `form:"title" json:"title" binding:"required"`
		Body  string `form:"body" json:"body" binding:"required"`
	}

	var json UpdatePost

	if err := c.ShouldBindJSON(&json); err == nil {
		editPost, err := db.Prepare("UPDATE post SET title=?, content=? WHERE id= " + c.Param("id"))
		if err != nil {
			c.JSON(500, gin.H{
				"messages": err,
			})
		}

		_, errEdit := editPost.Exec(json.Title, json.Body)
		if errEdit != nil {
			c.JSON(500, gin.H{
				"messages": errEdit,
			})
		} else {
			c.JSON(200, gin.H{
				"messages": "updated",
			})
		}
	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}

}

func Delete(c *gin.Context) {
	db := database.DBConn()

	delForm, err := db.Prepare("DELETE FROM post WHERE id=?")

	if err != nil {
		c.JSON(500, gin.H{
			"messages": err,
		})
	}

	delForm.Exec(c.Param("id"))
	c.JSON(200, gin.H{
		"messages": "deleted",
	})

	defer db.Close()
}
