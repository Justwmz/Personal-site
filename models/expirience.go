package models

import (
	"time"

	"gorm.io/gorm"
)

type Experience struct {
	gorm.Model
	Position string
	Company  string
	Summary  string
	From     time.Time
	To       time.Time
}
