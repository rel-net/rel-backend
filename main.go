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
	// CONTACT
	r.POST("/api/contact", controllers.CreateContact)
	r.GET("/api/contact", controllers.ListContacts)
	r.GET("/api/contact/:id", controllers.GetContact)
	r.PUT("/api/contact/:id", controllers.UpdateContact)
	r.DELETE("/api/contact/:id", controllers.DeleteContact)

	// NOTE
	r.POST("/api/note/contact/:user_id", controllers.CreateNote)
	r.GET("/api/note/contact/:user_id", controllers.ListNote)
	r.GET("/api/note/:note_id", controllers.GetNote)
	r.PUT("/api/note/:note_id", controllers.UpdateNote)
	r.DELETE("/api/note/:note_id", controllers.DeleteNote)
	r.Run()
}
