package services

import "github.com/JoaoRafa19/taskfy/internal/store"

type TaskService struct {
	Store store.TaskStore
}

func NewTaskService(store store.TaskStore) *TaskService {
	return &TaskService{
		Store: store,
	}
}

func (s *TaskService) CreateTask(title, description string, priotrity int32) (store.Task, error) {
	//validação e regra de negócios
	//

	task, err := s.Store.CreateTask(title, description, priotrity)

	if err != nil {
		return store.Task{}, err
	}

	return task, nil
}

func (s *TaskService) GetTaskById(id int32) (store.Task, error) {
	task, err := s.Store.GetTaskById(id)
	if err != nil {
		return store.Task{}, err
	}

	return task, nil
}

func (s *TaskService) UpdateTask(id int32, title, description string, priority int32) (store.Task, error) {

	task, err := s.Store.UpdateTask(id, title, description, priority)

	if err != nil {
		return store.Task{}, err
	}

	return task, nil
}

func (s *TaskService) DeleteTask(id int32) error {
	return s.Store.DeleteTask(id)
}
