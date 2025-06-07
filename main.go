package main

import (
	"log"
	"net/http"
	"task_manager_go/config"
	"task_manager_go/controller"
	repository2 "task_manager_go/repository"
	"task_manager_go/service"

	"github.com/gorilla/mux"
)

// main is the entry point of the application.
// Initializes the database connection, sets up the dependency chain,
// configures the router with REST API endpoints, and starts the HTTP server.
func main() {
	db := config.InitDB()
	repository := repository2.NewTaskRepository(db)
	taskService := service.NewTaskService(repository)
	taskController := controller.NewTaskController(taskService)
	r := mux.NewRouter()
	r.HandleFunc("/tasks", taskController.CreateTask).Methods("POST")
	r.HandleFunc("/tasks", taskController.GetAllTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", taskController.FindTaskById).Methods("GET")
	r.HandleFunc("/tasks/{id}", taskController.UpdateTaskById).Methods("PATCH")
	r.HandleFunc("/tasks/{id}", taskController.DeleteById).Methods("DELETE")

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Println("err with conn localhost ", err)
		return
	}
}
