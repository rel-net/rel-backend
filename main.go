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
	r.POST("/contact", controllers.CreateContact)
	r.GET("/contact", controllers.ListContacts)
	r.GET("/contact/:id", controllers.GetContact)
	r.PUT("/contact/:id", controllers.UpdateContact)
	r.DELETE("/contact/:id", controllers.DeleteContact)

	// NOTE
	r.POST("/note/:user_id", controllers.CreateNote)
	r.GET("/note/:user_id", controllers.ListNote)
	r.PUT("/note/:note_id", controllers.UpdateNote)
	r.DELETE("/note/:note_id", controllers.DeleteNote)
	r.Run()
}
