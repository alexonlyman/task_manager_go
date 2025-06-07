package service

import (
	"errors"
	"log"
	"task_manager_go/model"
	"task_manager_go/repository"
)

// TaskService provides business logic for task management.
// Uses TaskRepositoryInterface for data storage interaction.
type TaskService struct {
	repo repository.TaskRepositoryInterface
}

// NewTaskService creates a new instance of TaskService with the specified repository.
func NewTaskService(repo repository.TaskRepositoryInterface) *TaskService {
	return &TaskService{repo: repo}
}

// CreateTask creates a new task in the system.
// Returns the created task and an error if one occurred.
func (t *TaskService) CreateTask(task model.Task) (model.Task, error) {
	return t.repo.CreateTask(task)
}

// UpdateTask updates an existing task by its ID.
// Returns the updated task and an error if the task was not found or another error occurred.
func (t *TaskService) UpdateTask(id uint, task model.Task) (model.Task, error) {
	updatedTask, err := t.repo.FindById(id)
	if err != nil {
		return model.Task{}, err
	}
	if updatedTask.Id == 0 {
		log.Printf("Task with ID %d wasn't found", id)
		return model.Task{}, errors.New("task wasn't found")
	}
	return t.repo.UpdateTaskById(id, task)
}

// GetAllTasks returns a list of all tasks in the system.
// Returns a slice of tasks and an error if one occurred.
func (t *TaskService) GetAllTasks() ([]model.Task, error) {
	return t.repo.GetAll()
}

// GetTaskByID finds a task by its ID.
// Returns the found task and an error if the task was not found or another error occurred.
func (t *TaskService) GetTaskByID(id uint) (model.Task, error) {
	task, err := t.repo.FindById(id)
	if err != nil {
		return model.Task{}, err
	}
	if task.Id == 0 {
		log.Printf("Task with ID %d wasn't found", id)
		return model.Task{}, errors.New("task wasn't found")
	}
	return task, nil
}

// DeleteById deletes a task by its ID.
// Returns an error if the task was not found or another error occurred.
func (t *TaskService) DeleteById(id uint) error {
	task, err := t.repo.FindById(id)
	if err != nil {
		return err
	}
	if task.Id == 0 {
		log.Printf("Task with ID %d wasn't found", id)
		return errors.New("task wasn't found")
	}
	return t.repo.DeleteByID(id)
}
