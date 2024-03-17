package main

import (
	"fmt"
	"rel/controllers"
	"rel/initializers"
	"rel/middlewares"
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
	config.AllowOrigins = []string{"https://localhost:5173"} // Replace with your actual frontend domain
	config.AllowCredentials = true
	r.Use(cors.New(config))

	r.POST("/api/login", controllers.LoginHandler)
	r.POST("/api/signup", controllers.SignupHandler)
	r.GET("/api/validate", middlewares.RequireAuth, controllers.ValidateHandler) // here RequireAuth is a middleware that we will be creating below. It protects the route

	// USER
	r.POST("/api/user", controllers.CreateUser)
	r.GET("/api/user", controllers.ListUsers)
	r.GET("/api/user/:id", controllers.GetUser)
	r.PUT("/api/user/:id", controllers.UpdateUser)
	r.DELETE("/api/user/:id", controllers.DeleteUser)

	// CONTACT
	r.POST("/api/contact", middlewares.RequireAuth, controllers.CreateContact)
	r.GET("/api/contact", middlewares.RequireAuth, controllers.ListContacts)
	r.GET("/api/contact/:id", middlewares.RequireAuth, controllers.GetContact)
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

	r.RunTLS(":3000", "./certs/server.crt", "./certs/server.key")
}

func startReminderScheduler() {
	for {
		var reminders []models.Reminder
		initializers.DB.Find(&reminders)

		fmt.Println("Heartbeat: scheduler is alive")

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
