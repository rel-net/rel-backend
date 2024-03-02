package controllers

import (
	"rel/initializers"
	"rel/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var body struct {
		FirstName string
		LastName  string
		Email     string
		Phone     string
		LinkedIn  string
		Bio       string
	}

	c.Bind(&body)
	user := models.User{FirstName: body.FirstName, LastName: body.LastName, Email: body.Email, Phone: body.Phone, LinkedIn: body.LinkedIn, Bio: body.Bio}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"user": user,
	})
}

func ListUsers(c *gin.Context) {
	var users []models.User
	initializers.DB.Find(&users)

	c.JSON(200, gin.H{
		"users": users,
	})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	initializers.DB.First(&user, id)

	c.JSON(200, gin.H{
		"user": user,
	})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		FirstName string
		LastName  string
		Email     string
		Phone     string
		LinkedIn  string
		Bio       string
	}

	var user models.User
	initializers.DB.First(&user, id)

	c.Bind(&body)
	initializers.DB.Model(&user).Updates(models.User{FirstName: body.FirstName, LastName: body.LastName, Email: body.Email, Phone: body.Phone, LinkedIn: body.LinkedIn, Bio: body.Bio})

	c.JSON(200, gin.H{
		"user": user,
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	initializers.DB.Delete(&models.User{}, id)
	c.Status(200)
}
