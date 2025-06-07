package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"task_manager_go/model"
	"task_manager_go/service"

	"github.com/gorilla/mux"
)

// TaskController handles HTTP requests for task management.
// Provides REST API endpoints for CRUD operations on tasks.
type TaskController struct {
	service *service.TaskService
}

// NewTaskController creates a new instance of TaskController with the specified service.
func NewTaskController(service *service.TaskService) *TaskController {
	return &TaskController{service: service}
}

// GetAllTasks handles GET request to retrieve all tasks.
// Returns a JSON array of tasks or an error response.
func (c *TaskController) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := c.service.GetAllTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		http.Error(w, "Failed to encode tasks", http.StatusInternalServerError)
		return
	}
}

// CreateTask handles POST request to create a new task.
// Expects task data in JSON format in the request body.
// Returns the created task or an error response.
func (c *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	var createdTask model.Task
	err := json.NewDecoder(r.Body).Decode(&createdTask)
	if err != nil {
		http.Error(w, "Failed to decode tasks", http.StatusBadRequest)
		return
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Println("Failed to close request body:", err)
		}
	}()
	log.Printf("Received task data: %+v\n", createdTask)
	createdTaskPtr, err := c.service.CreateTask(createdTask)
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdTaskPtr)
}

// FindTaskById handles GET request to retrieve a task by its ID.
// Expects task ID in the URL path.
// Returns the found task or an error response.
func (c *TaskController) FindTaskById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Printf("Received ID: %v\n", vars["id"])

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("failed to convert variable")
		return
	}
	log.Printf("Parsed ID: %d\n", id)

	task, err := c.service.GetTaskByID(uint(id))
	if err != nil {
		log.Println("no task with id", id)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		http.Error(w, "Failed to encode task", http.StatusInternalServerError)
		return
	}
}

// UpdateTaskById handles PATCH request to update an existing task.
// Expects task ID in the URL path and updated task data in JSON format in the request body.
// Returns the updated task or an error response.
func (c *TaskController) UpdateTaskById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Printf("Received ID: %v\n", vars["id"])

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("failed to convert variable")
		return
	}
	var updatedTask model.Task

	log.Printf("Parsed ID: %d\n", id)

	err = json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		http.Error(w, "Failed to decode tasks", http.StatusBadRequest)
		return
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Println("Failed to close request body:", err)
		}
	}()

	w.Header().Set("Content-Type", "application/json")

	updatedTaskPtr, err := c.service.UpdateTask(uint(id), updatedTask)
	log.Println("task: ", updatedTask)
	if err != nil {
		http.Error(w, "failed with update task", http.StatusBadRequest)
		return
	}
	log.Println("update complete")
	json.NewEncoder(w).Encode(updatedTaskPtr)
}

// DeleteById handles DELETE request to remove a task.
// Expects task ID in the URL path.
// Returns success or error response.
func (c *TaskController) DeleteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Printf("Received ID: %v\n", vars["id"])

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("failed to convert variable")
		return
	}
	log.Printf("Parsed ID: %d\n", id)

	err = c.service.DeleteById(uint(id))
	if err != nil {
		log.Println("no task with id", id)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	log.Println("deleting complete")
}
