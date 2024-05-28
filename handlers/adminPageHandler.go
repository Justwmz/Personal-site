package handlers

import (
	"net/http"
	"time"

	"github.com/Justwmz/personal-site/components/admin"
	"github.com/Justwmz/personal-site/database"
	"github.com/Justwmz/personal-site/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
)

func IndexAdminPage(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	_ = claims

	var (
		certifications []models.Certifications
		experience     []models.Experience
		education      []models.Education
	)

	database.DB.Find(&certifications)

	database.DB.Find(&experience)

	database.DB.Find(&education)

	Render(w, r, admin.AdminIndex(certifications, experience, education))
}

func AddCert(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	provider := r.FormValue("provider")
	credentialID := r.FormValue("credentialID")
	link := r.FormValue("link")
	issued, _ := time.Parse("2006-01-02", r.FormValue("issued"))
	expired, _ := time.Parse("2006-01-02", r.FormValue("expired"))

	cert := models.Certifications{
		Name:         name,
		Provider:     provider,
		CredentialID: credentialID,
		Link:         link,
		Issued:       issued,
		Expired:      expired,
	}

	database.DB.Create(&cert)

	certs := []models.Certifications{}

	database.DB.Find(&certs)

	Render(w, r, admin.Certification(certs))
}

func AddExperience(w http.ResponseWriter, r *http.Request) {
	position := r.FormValue("position")
	company := r.FormValue("company")
	summary := r.FormValue("summary")
	from, _ := time.Parse("2006-01-02", r.FormValue("from"))
	to, _ := time.Parse("2006-01-02", r.FormValue("to"))

	exp := models.Experience{
		Position: position,
		Company:  company,
		Summary:  summary,
		From:     from,
		To:       to,
	}

	database.DB.Create(&exp)

	exps := []models.Experience{}

	database.DB.Find(&exps)

	Render(w, r, admin.WorkExperience(exps))
}

func AddEducation(w http.ResponseWriter, r *http.Request) {
	universityName := r.FormValue("universityName")
	degree := r.FormValue("degree")
	field := r.FormValue("field")
	from, _ := time.Parse("2006-01-02", r.FormValue("from"))
	to, _ := time.Parse("2006-01-02", r.FormValue("to"))

	education := models.Education{
		UniversityName: universityName,
		Degree:         degree,
		Field:          field,
		From:           from,
		To:             to,
	}

	database.DB.Create(&education)

	educations := []models.Education{}

	database.DB.Find(&educations)

	Render(w, r, admin.Education(educations))
}

func DeleteCert(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	cert := models.Certifications{}

	database.DB.Where("id = ?", id).Delete(&cert)

	certs := []models.Certifications{}

	database.DB.Find(&certs)

	Render(w, r, admin.Certification(certs))
}

func DeleteExperience(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	exp := models.Experience{}

	database.DB.Where("id = ?", id).Delete(&exp)

	exps := []models.Experience{}

	database.DB.Find(&exps)

	Render(w, r, admin.WorkExperience(exps))
}

func DeleteEducation(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	education := models.Education{}

	database.DB.Where("id = ?", id).Delete(&education)

	educations := []models.Education{}

	database.DB.Find(&educations)

	Render(w, r, admin.Education(educations))
}
