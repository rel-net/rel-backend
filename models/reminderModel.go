package models

import (
	"time"

	"gorm.io/gorm"
)

type Reminder struct {
	gorm.Model
	ContactId uint64
	Date      time.Time
	Todo      string
	Status    string
}
