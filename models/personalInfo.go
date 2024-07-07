package models

import "gorm.io/gorm"

type PersonalInfo struct {
	gorm.Model
	Country   string
	City      string
	FirstName string
	LastName  string
	Email     string
	Summary   string
}
