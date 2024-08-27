package routes

import (
	
	"backend/controllers"

	"github.com/gorilla/mux"
)

func UserRoutes()*mux.Router{
	r:= mux.NewRouter()
	r.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.GetUserById).Methods("GET")
	return r
}

