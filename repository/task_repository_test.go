package repository

import (
	"task_manager_go/config"
	"task_manager_go/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTaskRepository(t *testing.T) {
	db, cleanup := config.InitTestDBWithDocker()
	defer cleanup()
	repo := NewTaskRepository(db)
	assert.NotNil(t, repo)
	assert.Implements(t, (*TaskRepositoryInterface)(nil), repo)
}

func TestTaskRepository_CreateTask(t *testing.T) {
	db, cleanup := config.InitTestDBWithDocker()
	defer cleanup()
	repo := NewTaskRepository(db)
	task := model.Task{
		Id:     0,
		Name:   "name",
		Status: "status",
		Date:   time.Now().Truncate(time.Millisecond),
	}
	newTask, err := repo.CreateTask(task)
	if err != nil {
		return
	}
	assert.Equal(t, newTask.Name, "name")
	assert.NotNil(t, newTask)
}

func TestTaskRepository_GetAll(t *testing.T) {
	db, cleanup := config.InitTestDBWithDocker()
	defer cleanup()
	repo := NewTaskRepository(db)
	var tasks []model.Task
	task1 := model.Task{
		Id:     1,
		Name:   "name1",
		Status: "status1",
		Date:   time.Now().Truncate(time.Millisecond),
	}
	task2 := model.Task{
		Id:     2,
		Name:   "name2",
		Status: "status2",
		Date:   time.Now().Truncate(time.Millisecond),
	}
	tasks = append(tasks, task1, task2)

	_, err := repo.CreateTask(task1)
	assert.NoError(t, err)
	_, err = repo.CreateTask(task2)
	assert.NoError(t, err)

	allTasks, err := repo.GetAll()
	if err != nil {
		return
	}
	assert.NotNil(t, allTasks)
	assert.Equal(t, allTasks, tasks)
}
func TestTaskRepository_UpdateTaskById(t *testing.T) {
	db, cleanup := config.InitTestDBWithDocker()
	defer cleanup()
	repo := NewTaskRepository(db)
	task := model.Task{
		Id:     1,
		Name:   "name",
		Status: "status",
		Date:   time.Now().Truncate(time.Millisecond),
	}
	updatedTask := model.Task{
		Id:     1,
		Name:   "new_name",
		Status: "new_status",
		Date:   time.Now().Truncate(time.Millisecond),
	}
	createTask, err := repo.CreateTask(task)
	assert.NoError(t, err)

	foundTask, err := repo.FindById(task.Id)
	assert.NoError(t, err)

	updateTaskById, err := repo.UpdateTaskById(foundTask.Id, updatedTask)
	if err != nil {
		return
	}
	assert.NotEmpty(t, updateTaskById)
	assert.NotNil(t, foundTask)
	assert.NotEqual(t, createTask, updateTaskById)
	assert.Equal(t, foundTask.Name, "name")
	assert.Equal(t, updateTaskById.Name, "new_name")
}

func TestTaskRepository_FindById(t *testing.T) {
	db, cleanup := config.InitTestDBWithDocker()
	defer cleanup()
	repo := NewTaskRepository(db)

	task := model.Task{
		Id:     1,
		Name:   "test_task",
		Status: "test_status",
		Date:   time.Now().Truncate(time.Millisecond),
	}
	createdTask, err := repo.CreateTask(task)
	assert.NoError(t, err)
	assert.NotNil(t, createdTask)

	foundTask, err := repo.FindById(createdTask.Id)
	assert.NoError(t, err)
	assert.NotNil(t, foundTask)
	assert.Equal(t, createdTask.Name, foundTask.Name)
	assert.Equal(t, createdTask.Status, foundTask.Status)

	nonExistentTask, err := repo.FindById(999)
	assert.Error(t, err)
	assert.Empty(t, nonExistentTask.Id)
}

func TestTaskRepository_DeleteByID(t *testing.T) {
	db, cleanup := config.InitTestDBWithDocker()
	defer cleanup()
	repo := NewTaskRepository(db)

	task := model.Task{
		Id:     1,
		Name:   "task_to_delete",
		Status: "status_to_delete",
		Date:   time.Now().Truncate(time.Millisecond),
	}
	createdTask, err := repo.CreateTask(task)
	assert.NoError(t, err)
	assert.NotNil(t, createdTask)

	err = repo.DeleteByID(createdTask.Id)
	assert.NoError(t, err)

	deletedTask, err := repo.FindById(createdTask.Id)
	assert.Error(t, err)
	assert.Empty(t, deletedTask.Id)

	err = repo.DeleteByID(999)
	assert.Error(t, err, "Ожидается ошибка при удалении несуществующей задачи")
}
