package handlers

import (
	"net/http"

	"github.com/Justwmz/personal-site/components"
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	Render(w, r, components.Index())
}
