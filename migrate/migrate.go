package main

import (
	"rel/initializers"
	"rel/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Contact{})
}
