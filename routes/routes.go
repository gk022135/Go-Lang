package routes

import (
	"net/http"
	"go-prisma-backend/controllers/auths"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/signup", auths.SignUp)
	mux.HandleFunc("/login", auths.Login)
}
