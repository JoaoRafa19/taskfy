package services

import (
	"errors"
	"testing"
	"time"

	"github.com/JoaoRafa19/taskfy/internal/store"
	"github.com/stretchr/testify/assert"
)

type MockTaskStore struct {
}

func (s *MockTaskStore) CreateTask(title, description string, priotrity int32) (store.Task, error) {
	return store.Task{
		Id:          1,
		Title:       title,
		Description: description,
		Priority:    priotrity,
		CreatedAt:   time.Now(),
		UpdateAt:    time.Now(),
	}, nil
}

func (s *MockTaskStore) GetTask(id int32) (store.Task, error) {
	return store.Task{
		Id:          1,
		Title:       "Teste Task",
		Description: "description",
		Priority:    1,
		CreatedAt:   time.Now(),
		UpdateAt:    time.Now(),
	}, nil
}

func (s *MockTaskStore) UpdateTask(id int32, title, description string, priority int32) (store.Task, error) {

	return store.Task{
		Id:          1,
		Title:       "Teste Task",
		Description: "description",
		Priority:    1,
		CreatedAt:   time.Now().Add(-10 * time.Minute),
		UpdateAt:    time.Now(),
	}, nil
}

func (s *MockTaskStore) DeleteTask(id int32) error {
	return nil
}

func (s *MockTaskStore) ListTasks() ([]store.Task, error) {
	return []store.Task{
		store.Task{
			Id:          1,
			Title:       "Teste Task",
			Description: "description",
			Priority:    1,
			CreatedAt:   time.Now(),
			UpdateAt:    time.Now(),
		},
		store.Task{
			Id:          2,
			Title:       "Teste Task",
			Description: "description",
			Priority:    2,
			CreatedAt:   time.Now(),
			UpdateAt:    time.Now(),
		},
	}, nil
}

func (s *MockTaskStore) GetTaskById(id int32) (store.Task, error) {

	if id != 2 {
		return store.Task{}, errors.New("Not Found")
	}

	return store.Task{
		Id:          2,
		Title:       "Teste Task",
		Description: "description",
		Priority:    2,
		CreatedAt:   time.Now(),
		UpdateAt:    time.Now(),
	}, nil
}

func TestCreateTask(t *testing.T) {
	mockStore := &MockTaskStore{}

	taskService := NewTaskService(mockStore)

	task, err := taskService.CreateTask("Mock Test Task", "Mock Description", 1)

	if err != nil {
		t.Fatal(err)
	}

	assert.NoError(t, err)
	assert.Equal(t, "Mock Test Task", task.Title)
	assert.Equal(t, "Mock Description", task.Description)
	assert.Equal(t, int32(1), task.Priority)
}

func TestGetTask(t *testing.T) {
	mockStore := &MockTaskStore{}

	taskService := NewTaskService(mockStore)

	task, err := taskService.GetTaskById(2)
	assert.Nil(t, err)
	assert.Equal(t, int32(2), task.Id)
	assert.Equal(t, "Teste Task", task.Title)
}

func TestListTasks(t *testing.T) {
	mockStore := &MockTaskStore{}

	taskService := NewTaskService(mockStore)

	tasks, err := taskService.Store.ListTasks()
	assert.NoError(t, err)
	assert.Len(t, tasks, 2)
}
