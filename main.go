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
	r.POST("/contact", controllers.CreateContact)
	r.GET("/contact", controllers.ListContacts)
	r.GET("/contact/:id", controllers.GetContact)
	r.PUT("/contact/:id", controllers.UpdateContact)
	r.DELETE("/contact/:id", controllers.DeleteContact)
	r.Run()
}
