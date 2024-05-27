package main

import (
	"net/http"

	"github.com/Justwmz/personal-site/auth"
	"github.com/Justwmz/personal-site/database"
	"github.com/Justwmz/personal-site/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
)

func main() {
	database.InitDatabase()
	auth.InitiAuth()

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
		})
	})

	http.ListenAndServe(":3000", r)
}
