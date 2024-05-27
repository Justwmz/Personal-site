package models

import (
	"time"

	"gorm.io/gorm"
)

type Certifications struct {
	gorm.Model
	Name         string
	Provider     string
	CredentialID string
	Link         string
	Issued       time.Time
	Expired      time.Time
}
