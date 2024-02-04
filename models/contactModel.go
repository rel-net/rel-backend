package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Name     string
	LastName string
	Email    string
	Phone    string
	LinkedIn string
}
