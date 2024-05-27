package main

import (
	"net/http"

	"github.com/Justwmz/personal-site/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Handle("/*", public())
	r.Get("/", handlers.IndexPage)
	http.ListenAndServe(":3000", r)
}
