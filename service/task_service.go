package main

import (
	"errors"
	"log"
	"task_manager_go/model"
	"task_manager_go/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (t *TaskService) CreateTask(task model.Task) error {
	return t.repo.CreateTask(task)
}

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
func (t *TaskService) GetAllTasks() ([]model.Task, error) {
	return t.repo.GetAll()
}

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
