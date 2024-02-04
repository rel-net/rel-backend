package models

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	ContactId uint64
	Date      time.Time
	Content   string
}
