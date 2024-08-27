package routes

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	RegisterAuthRoutes(r)
	UserRoutes()
	return r
}
