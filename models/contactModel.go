package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	UserID             uint64 // ID of the user who created the contact (foreign key referencing User model)
	Name               string
	LastName           string
	Email              string
	Phone              string
	LinkedIn           string
	IsUser             bool
	ContactUserId      uint64 // ID of the user when isUser = true
	InvitationSent     bool
	InvitationAccepted bool
}
