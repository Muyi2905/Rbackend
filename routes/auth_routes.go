package routes

import (
	"backend/controllers"


	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(r *mux.Router){
	r.HandleFunc("/register", controllers.RegisterHandler).Methods("POST")

	r.HandleFunc("/login", controllers.LoginHandler).Methods("POST")
}
