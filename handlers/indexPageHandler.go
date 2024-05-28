package handlers

import (
	"net/http"

	"github.com/Justwmz/personal-site/components"
	"github.com/Justwmz/personal-site/database"
	"github.com/Justwmz/personal-site/models"
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	var (
		certifications []models.Certifications
		experience     []models.Experience
		education      []models.Education
	)

	database.DB.Find(&certifications)

	database.DB.Find(&experience)

	database.DB.Find(&education)

	Render(w, r, components.Index(certifications, education, experience))
}
