package controllers

import (
	"net/http"
	"rel/initializers"
	"rel/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateContact(c *gin.Context) {
	var body struct {
		Name     string
		LastName string
		Email    string
		Phone    string
		LinkedIn string
	}

	c.Bind(&body)
	contact := models.Contact{Name: body.Name, LastName: body.LastName, Email: body.Email, Phone: body.Phone, LinkedIn: body.LinkedIn}
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
	initializers.DB.Find(&contacts)

	if strings.HasPrefix(c.Request.URL.Path, "/home") {
		// Handle HTML rendering logic
		// You can load HTML templates and render them

		c.HTML(http.StatusOK, "contacts.html", gin.H{
			"contacts": contacts,
		})
		return
	}

	c.JSON(200, gin.H{
		"contacts": contacts,
	})
}

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
	return
}

func GetContact(c *gin.Context) {
	id := c.Param("contact_id")

	var contact models.Contact
	initializers.DB.First(&contact, id)

	if strings.HasPrefix(c.Request.URL.Path, "/home") {
		// Handle HTML rendering logic
		// You can load HTML templates and render them

		c.HTML(http.StatusOK, "test.html", gin.H{
			"contact": contact,
		})
		return
	}

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
	}

	var contact models.Contact
	initializers.DB.First(&contact, id)

	c.Bind(&body)
	initializers.DB.Model(&contact).Updates(models.Contact{Name: body.Name, LastName: body.LastName, Email: body.Email, Phone: body.Phone, LinkedIn: body.LinkedIn})

	c.JSON(200, gin.H{
		"contact": contact,
	})
}

func DeleteContact(c *gin.Context) {
	id := c.Param("id")
	initializers.DB.Delete(&models.Contact{}, id)
	c.Status(200)
}
