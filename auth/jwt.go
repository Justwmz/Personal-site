package auth

import (
	"github.com/go-chi/jwtauth"
)

var TokenAuth *jwtauth.JWTAuth

func InitiAuth() {
	TokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
}
