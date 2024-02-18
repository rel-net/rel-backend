package controllers

import (
	"rel/initializers"
	"rel/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateReminder(c *gin.Context) {

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
		Date   time.Time
		Todo   string
		Status string
	}

	c.Bind(&body)

	reminder := models.Reminder{
		ContactId: contactIdUIint,
		Date:      body.Date,
		Todo:      body.Todo,
		Status:    "Pending",
	}

	result := initializers.DB.Create(&reminder)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"reminder": reminder,
	})
}

func ListReminder(c *gin.Context) {
	var reminders []models.Reminder
	initializers.DB.Find(&reminders)

	c.JSON(200, gin.H{
		"reminders": reminders,
	})
}

func GetReminder(c *gin.Context) {
	id := c.Param("reminder_id")

	var reminder models.Reminder
	initializers.DB.First(&reminder, id)

	c.JSON(200, gin.H{
		"reminder": reminder,
	})
}

func ListContactReminder(c *gin.Context) {
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

	var reminders []models.Reminder
	if err := initializers.DB.Where("contact_id = ?", contactIdUIint).Find(&reminders).Error; err != nil {
		c.Status(500)
		return
	}

	c.JSON(200, gin.H{
		"reminders": reminders,
	})
}

func UpdateReminder(c *gin.Context) {
	id := c.Param("reminder_id")

	var body struct {
		ContactId uint64
		Date      time.Time
		Todo      string
		Status    string
	}

	var reminder models.Reminder
	initializers.DB.First(&reminder, id)

	c.Bind(&body)
	initializers.DB.Model(&reminder).Updates(models.Reminder{ContactId: body.ContactId, Date: body.Date, Todo: body.Todo, Status: body.Status})

	c.JSON(200, gin.H{
		"reminder": reminder,
	})
}

func DeleteReminder(c *gin.Context) {
	id := c.Param("reminder_id")
	initializers.DB.Delete(&models.Reminder{}, id)
	c.Status(200)
}
