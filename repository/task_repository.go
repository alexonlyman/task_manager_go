package repository

import (
	"task_manager_go/model"

	"gorm.io/gorm"
)

// TaskRepositoryInterface defines the contract for task data storage operations.
type TaskRepositoryInterface interface {
	// CreateTask stores a new task in the database.
	CreateTask(task model.Task) (model.Task, error)
	// GetAll retrieves all tasks from the database.
	GetAll() ([]model.Task, error)
	// FindById retrieves a task by its ID from the database.
	FindById(id uint) (model.Task, error)
	// UpdateTaskById updates an existing task in the database.
	UpdateTaskById(id uint, task model.Task) (model.Task, error)
	// DeleteByID removes a task from the database by its ID.
	DeleteByID(id uint) error
}

// TaskRepository implements TaskRepositoryInterface using GORM for database operations.
type TaskRepository struct {
	db *gorm.DB
}

// NewTaskRepository creates a new instance of TaskRepository with the specified database connection.
func NewTaskRepository(db *gorm.DB) TaskRepositoryInterface {
	return &TaskRepository{db: db}
}

// CreateTask implements the creation of a new task in the database.
func (r *TaskRepository) CreateTask(task model.Task) (model.Task, error) {
	result := r.db.Create(&task)
	return task, result.Error
}

// GetAll implements the retrieval of all tasks from the database.
func (r *TaskRepository) GetAll() ([]model.Task, error) {
	var tasks []model.Task
	result := r.db.Find(&tasks)
	return tasks, result.Error
}

// FindById implements the retrieval of a task by its ID from the database.
func (r *TaskRepository) FindById(id uint) (model.Task, error) {
	var task model.Task
	result := r.db.First(&task, id)
	return task, result.Error
}

// UpdateTaskById implements the update of an existing task in the database.
func (r *TaskRepository) UpdateTaskById(id uint, task model.Task) (model.Task, error) {
	var updatedTask model.Task
	result := r.db.Model(&updatedTask).Where("Id = ?", id).Updates(map[string]interface{}{
		"name":   task.Name,
		"status": task.Status,
	}).First(&updatedTask)
	return updatedTask, result.Error
}

// DeleteByID implements the removal of a task from the database by its ID.
func (r *TaskRepository) DeleteByID(id uint) error {
	var task model.Task
	if err := r.db.First(&task, id).Error; err != nil {
		return err
	}
	result := r.db.Delete(&task)
	return result.Error
}
