package main

import (
	"fmt"
	"rel/controllers"
	"rel/initializers"
	"rel/models"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	// Use CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"} // Replace with your actual frontend domain
	r.Use(cors.New(config))

	// CONTACT
	r.POST("/api/contact", controllers.CreateContact)
	r.GET("/api/contact", controllers.ListContacts)
	r.GET("/api/contact/:id", controllers.GetContact)
	r.PUT("/api/contact/:id", controllers.UpdateContact)
	r.DELETE("/api/contact/:id", controllers.DeleteContact)

	// NOTE
	r.POST("/api/note/contact/:contact_id", controllers.CreateNote)
	r.GET("/api/note/contact/:contact_id", controllers.ListNote)
	r.GET("/api/note/:note_id", controllers.GetNote)
	r.PUT("/api/note/:note_id", controllers.UpdateNote)
	r.DELETE("/api/note/:note_id", controllers.DeleteNote)

	// REMINDER
	r.POST("/api/reminder/contact/:contact_id", controllers.CreateReminder)
	r.GET("/api/reminder/contact", controllers.ListReminder)
	r.GET("/api/reminder/contact/:contact_id", controllers.ListContactReminder)
	r.GET("/api/reminder/:reminder_id", controllers.GetReminder)
	r.PUT("/api/reminder/:reminder_id", controllers.UpdateReminder)
	r.DELETE("/api/reminder/:reminder_id", controllers.DeleteReminder)

	go startReminderScheduler()

	r.Run()
}

func startReminderScheduler() {
	for {
		var reminders []models.Reminder
		initializers.DB.Find(&reminders)

		fmt.Println("Hello from the scheduler")

		for _, reminder := range reminders {
			if reminder.Status == "Pending" {
				if reminder.Date.Before(time.Now()) || reminder.Date.Equal(time.Now()) {
					fmt.Printf("Reminder for Contact %d: %s\n %s\n", reminder.ContactId, reminder.Todo, reminder.Status)
					initializers.DB.Model(&reminder).Updates(models.Reminder{Status: "Sent"})
				}
			}

		}
		// Sleep for a certain interval before checking again
		time.Sleep(time.Minute)

	}
}
