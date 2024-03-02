package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName      string
	LastName       string
	Email          string
	LinkedIn       string
	Phone          string
	Password       string
	Bio            string
	ProfilePicture string
	JoinedAt       time.Time
}
