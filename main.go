package main

import (
	"net/http"

	"github.com/Justwmz/personal-site/auth"
	"github.com/Justwmz/personal-site/database"
	"github.com/Justwmz/personal-site/handlers"
	"github.com/Justwmz/personal-site/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
)

func main() {
	database.InitDatabase()
	auth.InitAuth()
	utils.GetCountryData()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Handle("/*", public())
	r.Get("/", handlers.IndexPage)

	r.Route("/admin", func(r chi.Router) {
		r.Get("/login", handlers.LoginAdminPage)
		r.Post("/login", handlers.LoginUser)

		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(auth.TokenAuth))
			r.Use(jwtauth.Authenticator)
			r.Get("/", handlers.IndexAdminPage)

			r.Get("/cert/{id}", handlers.GetCert)
			r.Get("/cert/{id}/edit", handlers.EditCert)
			r.Post("/cert", handlers.AddCert)
			r.Put("/cert/{id}", handlers.UpdateCert)
			r.Delete("/cert/{id}", handlers.DeleteCert)

			r.Get("/exp/{id}", handlers.GetWork)
			r.Get("/exp/{id}/edit", handlers.EditWork)
			r.Put("/exp/{id}", handlers.UpdateWork)
			r.Post("/exp", handlers.AddExperience)
			r.Delete("/exp/{id}", handlers.DeleteExperience)

			r.Get("/edu/{id}", handlers.GetEdu)
			r.Get("/edu/{id}/edit", handlers.EditEdu)
			r.Put("/edu/{id}", handlers.UpdateEdu)
			r.Post("/edu", handlers.AddEducation)
			r.Delete("/edu/{id}", handlers.DeleteEducation)

			r.Get("/personalInfo", handlers.GetPersonalInfo)
			r.Put("/personalInfo", handlers.AddPersonalInfo)
		})
	})

	http.ListenAndServe(":3000", r)
}
