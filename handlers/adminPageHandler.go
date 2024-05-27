package handlers

import (
	"net/http"

	"github.com/Justwmz/personal-site/components/admin"
	"github.com/go-chi/jwtauth"
)

func IndexAdminPage(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	_ = claims
	Render(w, r, admin.AdminIndex())
}
