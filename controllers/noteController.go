package controllers

import (
	"rel/initializers"
	"rel/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateNote(c *gin.Context) {

	contactId := c.Param("contact_id")
	contactIdUIint, err := strconv.ParseUint(contactId, 10, 64)
	if err != nil {
		c.Status(400)
		return
	}
	var contact models.Contact

	if err := initializers.DB.First(&contact, contactIdUIint).Error; err != nil {
		c.Status(404)
		return
	}

	var body struct {
		Content string
		Title   string
	}

	c.Bind(&body)

	note := models.Note{
		ContactId: contactIdUIint,
		Date:      time.Now(),
		Content:   body.Content,
		Title:     body.Title,
	}

	result := initializers.DB.Create(&note)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"note": note,
	})
}

func ListNote(c *gin.Context) {

	contactId := c.Param("contact_id")
	contactIdUIint, err := strconv.ParseUint(contactId, 10, 64)
	if err != nil {
		c.Status(400)
		return
	}
	// check if the user exists
	var contact models.Contact
	if err := initializers.DB.First(&contact, contactIdUIint).Error; err != nil {
		c.Status(404)
		return
	}

	var notes []models.Note
	if err := initializers.DB.Where("contact_id = ?", contactIdUIint).Order("date DESC").Find(&notes).Error; err != nil {
		c.Status(500)
		return
	}

	c.JSON(200, gin.H{
		"notes": notes,
	})
}

func GetNote(c *gin.Context) {
	id := c.Param("note_id")

	var note models.Note
	initializers.DB.First(&note, id)

	c.JSON(200, gin.H{
		"note": note,
	})
}

func UpdateNote(c *gin.Context) {
	id := c.Param("note_id")

	var body struct {
		ContactId uint64
		Date      time.Time
		Content   string
		Title     string
	}

	var note models.Note
	initializers.DB.First(&note, id)

	c.Bind(&body)
	initializers.DB.Model(&note).Updates(models.Note{ContactId: body.ContactId, Date: body.Date, Content: body.Content, Title: body.Title})

	c.JSON(200, gin.H{
		"note": note,
	})
}

func DeleteNote(c *gin.Context) {
	id := c.Param("note_id")
	initializers.DB.Delete(&models.Note{}, id)
	c.Status(200)
}
