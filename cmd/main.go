package main

import (
	"fmt"
	"log"
	"net/http"

	"backend/config"
	"backend/routes"
	"github.com/gorilla/mux"
)

func main() {

	err := config.InitDB()
	if err != nil {
		log.Fatalf("error connecting to database : %v", err)

	}

	r := mux.NewRouter()
	routes.RegisterRoutes()

	fmt.Println("server is starting on port 8080")

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
