package controllers

import (
	"fmt"
	"net/http"
	"rel/initializers"
	"rel/models"

	"github.com/gin-gonic/gin"
)

func CreateContact(c *gin.Context) {
	var body struct {
		UserID         uint64
		Name           string
		LastName       string
		Email          string
		Phone          string
		LinkedIn       string
		IsUser         bool
		InvitationSent bool
		ContactUserId  uint64
		Group          string
	}

	c.Bind(&body)
	contact := models.Contact{UserID: body.UserID, Name: body.Name, LastName: body.LastName, Email: body.Email, Phone: body.Phone, LinkedIn: body.LinkedIn, IsUser: body.IsUser, InvitationSent: body.InvitationSent, ContactUserId: body.ContactUserId, Group: body.Group}
	result := initializers.DB.Create(&contact)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"contact": contact,
	})
}

func ListContacts(c *gin.Context) {
	var contacts []models.Contact

	userInterface, exists := c.Get("user")
	if !exists {
		// Handle the case where the user is not found in the context
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, ok := userInterface.(models.User)
	if !ok {
		// Handle the case where the value stored is not of type models.User
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	initializers.DB.Where("user_id = ?", user.ID).Find(&contacts)
	fmt.Println(contacts)
	c.JSON(200, gin.H{
		"contacts": contacts,
	})
}

func GetContact(c *gin.Context) {
	id := c.Param("id")

	var contact models.Contact
	initializers.DB.First(&contact, id)

	c.JSON(200, gin.H{
		"contact": contact,
	})
}

func UpdateContact(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Name     string
		LastName string
		Email    string
		Phone    string
		LinkedIn string
		Group    string
	}

	var contact models.Contact
	initializers.DB.First(&contact, id)

	c.Bind(&body)
	initializers.DB.Model(&contact).Updates(models.Contact{Name: body.Name, LastName: body.LastName, Email: body.Email, Phone: body.Phone, LinkedIn: body.LinkedIn, Group: body.Group})

	c.JSON(200, gin.H{
		"contact": contact,
	})
}

func DeleteContact(c *gin.Context) {
	id := c.Param("id")
	initializers.DB.Delete(&models.Contact{}, id)
	c.Status(200)
}
