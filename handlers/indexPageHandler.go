package handlers

import (
	"net/http"

	"github.com/Justwmz/personal-site/components"
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	// var (
	// 	certifications []models.Certifications
	// 	experience     []models.Experience
	// 	education      []models.Education
	// )

	// certResult := database.DB.Find(&certifications)

	// expResult := database.DB.Find(&experience)

	// eduResult := database.DB.Find(&education)

	Render(w, r, components.Index())
}
