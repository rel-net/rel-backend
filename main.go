package main

import (
	"rel/controllers"
	"rel/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	r.GET("/home", controllers.Home)
	r.GET("/home/contact/:contact_id", controllers.GetContact)
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
	r.GET("/api/reminder/contact/:contact_id", controllers.ListReminder)
	r.GET("/api/reminder/:reminder_id", controllers.GetReminder)
	r.PUT("/api/reminder/:reminder_id", controllers.UpdateReminder)
	r.DELETE("/api/reminder/:reminder_id", controllers.DeleteReminder)
	r.Run()
}
