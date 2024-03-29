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
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Contact{})
	initializers.DB.AutoMigrate(&models.Note{})
	initializers.DB.AutoMigrate(&models.Reminder{})
}
