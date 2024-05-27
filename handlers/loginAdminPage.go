package handlers

import (
	"fmt"
	"net/http"

	"github.com/Justwmz/personal-site/auth"
	"github.com/Justwmz/personal-site/components/admin"
	"github.com/Justwmz/personal-site/database"
	"github.com/Justwmz/personal-site/models"
)

func LoginAdminPage(w http.ResponseWriter, r *http.Request) {
	Render(w, r, admin.AdminLogin(""))
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	user := models.User{}

	result := database.DB.Where("username = ?", username).Where("password = ?", password).First(&user)

	if result.RowsAffected > 0 {
		_, tokenString, _ := auth.TokenAuth.Encode(map[string]interface{}{"user_id": 1})
		fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)

		cookie := &http.Cookie{
			Name:     "jwt",
			Value:    tokenString,
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
			MaxAge:   3600,
		}

		http.SetCookie(w, cookie)
		w.Header().Set("HX-Redirect", "/admin")
	} else {
		w.Write([]byte("Something went wrong!"))
	}
}
