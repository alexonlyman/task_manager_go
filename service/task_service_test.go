package service

import (
	"errors"
	"task_manager_go/model"
	"task_manager_go/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTaskService_CreateTask(t *testing.T) {
	mockRepo := repository.NewMockTaskRepository()
	taskService := NewTaskService(mockRepo)

	task := model.Task{Name: "Test Task", Status: "Pending", Date: time.Now()}
	createdTask, err := taskService.CreateTask(task)

	assert.NoError(t, err)
	assert.NotZero(t, createdTask.Id)
	assert.Equal(t, "Test Task", createdTask.Name)
}

func TestTaskService_GetTaskByID(t *testing.T) {
	mockRepo := repository.NewMockTaskRepository()
	taskService := NewTaskService(mockRepo)

	task := model.Task{Name: "Test Task", Status: "Pending", Date: time.Now()}
	createdTask, _ := mockRepo.CreateTask(task)

	foundTask, err := taskService.GetTaskByID(createdTask.Id)
	assert.NoError(t, err)
	assert.Equal(t, createdTask.Id, foundTask.Id)

	_, err = taskService.GetTaskByID(999)
	assert.Error(t, err)
	assert.Equal(t, errors.New("task not found"), err)
}

func TestTaskService_UpdateTask(t *testing.T) {
	mockRepo := repository.NewMockTaskRepository()
	taskService := NewTaskService(mockRepo)

	task := model.Task{Name: "Original Task", Status: "Pending", Date: time.Now()}
	createdTask, _ := mockRepo.CreateTask(task)

	updatedTask := model.Task{Name: "Updated Task", Status: "Completed"}
	result, err := taskService.UpdateTask(createdTask.Id, updatedTask)

	assert.NoError(t, err)
	assert.Equal(t, "Updated Task", result.Name)
	assert.Equal(t, "Completed", result.Status)
}

func TestTaskService_DeleteById(t *testing.T) {
	mockRepo := repository.NewMockTaskRepository()
	taskService := NewTaskService(mockRepo)

	task := model.Task{Name: "Test Task", Status: "Pending", Date: time.Now()}
	createdTask, _ := mockRepo.CreateTask(task)

	err := taskService.DeleteById(createdTask.Id)
	assert.NoError(t, err)

	_, err = taskService.GetTaskByID(createdTask.Id)
	assert.Error(t, err)

	err = taskService.DeleteById(999)
	assert.Error(t, err)
}
