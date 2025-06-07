package test

import (
	"task_manager_go/config"
	"task_manager_go/model"
	"task_manager_go/repository"
	"task_manager_go/service"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func setupTestEnvironment(t *testing.T) (*service.TaskService, func()) {
	db, cleanup := config.InitTestDBWithDocker()
	repo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(repo)
	return taskService, cleanup
}

func createTestTask(t *testing.T, service *service.TaskService) model.Task {
	task := model.Task{
		Name:   "Test Task",
		Status: "Pending",
		Date:   time.Now().Truncate(time.Millisecond),
	}
	createdTask, err := service.CreateTask(task)
	assert.NoError(t, err)
	assert.NotNil(t, createdTask)
	return createdTask
}

func TestTaskCRUDIntegration(t *testing.T) {
	service, cleanup := setupTestEnvironment(t)
	defer cleanup()

	// Test Create
	t.Run("Create Task", func(t *testing.T) {
		task := model.Task{
			Name:   "Integration Test Task",
			Status: "New",
			Date:   time.Now().Truncate(time.Millisecond),
		}
		createdTask, err := service.CreateTask(task)
		assert.NoError(t, err)
		assert.NotNil(t, createdTask)
		assert.NotZero(t, createdTask.Id)
		assert.Equal(t, task.Name, createdTask.Name)
		assert.Equal(t, task.Status, createdTask.Status)
	})

	// Test Read
	t.Run("Get Task", func(t *testing.T) {
		createdTask := createTestTask(t, service)

		// Test Get by ID
		foundTask, err := service.GetTaskByID(createdTask.Id)
		assert.NoError(t, err)
		assert.Equal(t, createdTask.Id, foundTask.Id)
		assert.Equal(t, createdTask.Name, foundTask.Name)

		// Test Get All
		allTasks, err := service.GetAllTasks()
		assert.NoError(t, err)
		assert.NotEmpty(t, allTasks)
	})

	// Test Update
	t.Run("Update Task", func(t *testing.T) {
		createdTask := createTestTask(t, service)

		updatedTask := model.Task{
			Name:   "Updated Task",
			Status: "Completed",
			Date:   time.Now().Truncate(time.Millisecond),
		}

		result, err := service.UpdateTask(createdTask.Id, updatedTask)
		assert.NoError(t, err)
		assert.Equal(t, updatedTask.Name, result.Name)
		assert.Equal(t, updatedTask.Status, result.Status)
	})

	// Test Delete
	t.Run("Delete Task", func(t *testing.T) {
		createdTask := createTestTask(t, service)

		// Test successful deletion
		err := service.DeleteById(createdTask.Id)
		assert.NoError(t, err)

		// Verify task is deleted
		_, err = service.GetTaskByID(createdTask.Id)
		assert.Error(t, err)
	})
}

func TestTaskErrorCases(t *testing.T) {
	service, cleanup := setupTestEnvironment(t)
	defer cleanup()

	t.Run("Get Non-existent Task", func(t *testing.T) {
		_, err := service.GetTaskByID(999)
		assert.Error(t, err)
	})

	t.Run("Update Non-existent Task", func(t *testing.T) {
		nonExistentTask := model.Task{
			Name:   "Non-existent",
			Status: "Unknown",
		}
		_, err := service.UpdateTask(999, nonExistentTask)
		assert.Error(t, err)
	})

	t.Run("Delete Non-existent Task", func(t *testing.T) {
		err := service.DeleteById(999)
		assert.Error(t, err)
	})
}

func TestTaskConcurrentOperations(t *testing.T) {
	service, cleanup := setupTestEnvironment(t)
	defer cleanup()

	// Create multiple tasks concurrently
	tasks := make([]model.Task, 5)
	for i := range tasks {
		tasks[i] = model.Task{
			Name:   "Concurrent Task",
			Status: "Pending",
			Date:   time.Now().Truncate(time.Millisecond),
		}
	}

	// Test concurrent creation
	t.Run("Concurrent Creation", func(t *testing.T) {
		results := make(chan error, len(tasks))
		for _, task := range tasks {
			go func(t model.Task) {
				_, err := service.CreateTask(t)
				results <- err
			}(task)
		}

		for i := 0; i < len(tasks); i++ {
			err := <-results
			assert.NoError(t, err)
		}
	})
}
