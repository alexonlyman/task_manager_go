package repository

import (
	"errors"
	"task_manager_go/model"
)

type MockTaskRepository struct {
	tasks map[uint]model.Task
}

func NewMockTaskRepository() *MockTaskRepository {
	return &MockTaskRepository{
		tasks: make(map[uint]model.Task),
	}
}

func (m *MockTaskRepository) CreateTask(task model.Task) (model.Task, error) {
	task.Id = uint(len(m.tasks) + 1)
	m.tasks[task.Id] = task
	return task, nil
}

func (m *MockTaskRepository) GetAll() ([]model.Task, error) {
	tasks := make([]model.Task, 0, len(m.tasks))
	for _, task := range m.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (m *MockTaskRepository) FindById(id uint) (model.Task, error) {
	if task, exists := m.tasks[id]; exists {
		return task, nil
	}
	return model.Task{}, errors.New("task not found")
}

func (m *MockTaskRepository) UpdateTaskById(id uint, task model.Task) (model.Task, error) {
	if _, exists := m.tasks[id]; !exists {
		return model.Task{}, errors.New("task not found")
	}
	task.Id = id
	m.tasks[id] = task
	return task, nil
}

func (m *MockTaskRepository) DeleteByID(id uint) error {
	if _, exists := m.tasks[id]; !exists {
		return errors.New("task not found")
	}
	delete(m.tasks, id)
	return nil
}
