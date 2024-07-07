package database

import (
	"log"

	"github.com/Justwmz/personal-site/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	DB  *gorm.DB
	err error
)

func InitDatabase() {
	DB, err = gorm.Open(sqlite.Open("site.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %s", err)
	}

	DB.AutoMigrate(
		&models.Post{},
		&models.Experience{},
		&models.Certifications{},
		&models.Education{},
		&models.User{},
		&models.PersonalInfo{},
	)
	if err != nil {
		log.Fatalf("failed to migrate models %s", err)
	}

	// Initial credentials, which is need to be changed later on
	initialAdminUser := models.User{
		Username: "admin",
		Password: "admin",
	}

	DB.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&initialAdminUser)
}
