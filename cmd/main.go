package main

import (
	"net/http"

	"github.com/backend/pkg/api"
	"github.com/backend/pkg/models"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// Add new routers
	models.CreateDatabase()
	// models.AddContentsToDB()
	router.HandleFunc("/tasks", api.GetAllTasks).Methods("GET")
	router.HandleFunc("/task/{name}", api.GetTaskWithID).Methods("GET")
	http.ListenAndServe(":5000", router)
}
