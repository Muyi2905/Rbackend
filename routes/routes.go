package routes

import (
	"backend/routes"

	"github.com/gorilla/mux"
)

func RegisterRoutes () *mux.Router{
	r:= mux.NewRouter()

	routes.RegisterAuthRoutes(r)

	return r
}