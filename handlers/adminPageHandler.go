package handlers

import (
	"net/http"
	"time"

	"github.com/Justwmz/personal-site/components/admin"
	"github.com/Justwmz/personal-site/database"
	"github.com/Justwmz/personal-site/models"
	"github.com/Justwmz/personal-site/utils"

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
		personalInfo   models.PersonalInfo
	)

	database.DB.Find(&certifications)

	database.DB.Find(&experience)

	database.DB.Find(&education)

	database.DB.Last(&personalInfo)

	utils.Render(w, r, admin.AdminIndex(admin.AdminData{CertData: certifications, WorkExpData: experience, EduData: education, PersonalInfo: personalInfo}))
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

	utils.Render(w, r, admin.Certification(certs, admin.Alert{Message: "The certification was added successfully"}))
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

	utils.Render(w, r, admin.WorkExperience(exps, admin.Alert{Message: "The experience was added successfully"}))
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

	utils.Render(w, r, admin.Education(educations, admin.Alert{Message: "The education was added successfully"}))
}

func DeleteCert(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	cert := models.Certifications{}

	database.DB.Where("id = ?", id).Delete(&cert)

	w.WriteHeader(http.StatusOK)
}

func DeleteExperience(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	exp := models.Experience{}

	database.DB.Where("id = ?", id).Delete(&exp)

	w.WriteHeader(http.StatusOK)
}

func DeleteEducation(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	education := models.Education{}

	database.DB.Where("id = ?", id).Delete(&education)

	w.WriteHeader(http.StatusOK)
}

func GetCert(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	cert := models.Certifications{}

	database.DB.Where("id = ?", id).Find(&cert)

	utils.Render(w, r, admin.ListItemCert(cert))
}

func EditCert(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	cert := models.Certifications{}

	database.DB.Where("id = ?", id).Find(&cert)

	utils.Render(w, r, admin.EditFormCert(cert))
}

func UpdateCert(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
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

	database.DB.Model(&cert).Where("id = ?", id).Updates(&cert)
	database.DB.Where("id = ?", id).Find(&cert)

	utils.Render(w, r, admin.ListItemCert(cert))
}

func GetEdu(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	edu := models.Education{}

	database.DB.Where("id = ?", id).Find(&edu)

	utils.Render(w, r, admin.ListItemEdu(edu))
}

func EditEdu(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	edu := models.Education{}

	database.DB.Where("id = ?", id).Find(&edu)

	utils.Render(w, r, admin.EditFormEdu(edu))
}

func UpdateEdu(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	universityName := r.FormValue("universityName")
	degree := r.FormValue("degree")
	field := r.FormValue("field")
	from, _ := time.Parse("2006-01-02", r.FormValue("from"))
	to, _ := time.Parse("2006-01-02", r.FormValue("to"))

	edu := models.Education{
		UniversityName: universityName,
		Degree:         degree,
		Field:          field,
		From:           from,
		To:             to,
	}

	database.DB.Model(&edu).Where("id = ?", id).Updates(&edu)
	database.DB.Where("id = ?", id).Find(&edu)

	utils.Render(w, r, admin.ListItemEdu(edu))
}

func GetWork(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	exp := models.Experience{}

	database.DB.Where("id = ?", id).Find(&exp)

	utils.Render(w, r, admin.ListItemWork(exp))
}

func EditWork(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	exp := models.Experience{}

	database.DB.Where("id = ?", id).Find(&exp)

	utils.Render(w, r, admin.EditFormWork(exp))
}

func UpdateWork(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
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

	database.DB.Model(&exp).Where("id = ?", id).Updates(&exp)
	database.DB.Where("id = ?", id).Find(&exp)

	utils.Render(w, r, admin.ListItemWork(exp))
}

func AddPersonalInfo(w http.ResponseWriter, r *http.Request) {
	var personalInfo models.PersonalInfo

	country := r.FormValue("country")
	city := r.FormValue("city")
	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	email := r.FormValue("email")
	summary := r.FormValue("summary")

	database.DB.Create(&models.PersonalInfo{
		Country:   country,
		City:      city,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Summary:   summary,
	})

	database.DB.Last(&personalInfo)

	utils.Render(w, r, admin.PersonalInfo(personalInfo, admin.Alert{Message: "The personal info was added successfully"}))
}

func GetPersonalInfo(w http.ResponseWriter, r *http.Request) {
	var personalInfo models.PersonalInfo
	database.DB.Last(&personalInfo)
	utils.Render(w, r, admin.PersonalInfo(personalInfo, admin.Alert{Message: ""}))
}
