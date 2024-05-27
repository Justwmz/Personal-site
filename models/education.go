package models

import (
	"time"

	"gorm.io/gorm"
)

type Education struct {
	gorm.Model
	UniversityName string
	Degree         string
	Field          string
	From           time.Time
	To             time.Time
}
