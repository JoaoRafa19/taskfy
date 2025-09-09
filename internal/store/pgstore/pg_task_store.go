package pgstore

import (
	"context"

	"github.com/JoaoRafa19/taskfy/internal/store"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PGTaskStore struct {
	Queries *Queries
	Pool    *pgxpool.Pool
}

func NewPGTaskStore(pool *pgxpool.Pool) *PGTaskStore {
	return &PGTaskStore{
		Queries: New(pool),
		Pool:    pool,
	}
}

func (s *PGTaskStore) CreateTask(ctx context.Context, title, description string, priority int32) (store.Task, error) {
	result, err := s.Queries.CreateTask(ctx, CreateTaskParams{
		Title:       title,
		Description: description,
		Priority:    priority,
	})
	if err != nil {
		return store.Task{}, err
	}

	return store.Task{
		Id:          result.ID,
		Title:       result.Title,
		Priority:    result.Priority,
		Description: result.Description,
		CreatedAt:   result.CreatedAt.Time,
		UpdateAt:    result.UpdatedAt.Time,
	}, nil
}
func (s *PGTaskStore) GetTaskById(ctx context.Context, id int32) (store.Task, error) {
	result, err := s.Queries.GetTaskById(ctx, id)
	if err != nil {
		return store.Task{}, err
	}

	return store.Task{
		Id:          result.ID,
		Title:       result.Title,
		Priority:    result.Priority,
		Description: result.Description,
		CreatedAt:   result.CreatedAt.Time,
		UpdateAt:    result.UpdatedAt.Time,
	}, nil
}

func (s *PGTaskStore) ListTasks(ctx context.Context) ([]store.Task, error) {
	results, err := s.Queries.ListTasks(ctx)
	if err != nil {
		return nil, err
	}

	tasks := make([]store.Task, len(results))
	for i, result := range results {
		tasks[i] = store.Task{
			Id:          result.ID,
			Title:       result.Title,
			Priority:    result.Priority,
			Description: result.Description,
			CreatedAt:   result.CreatedAt.Time,
			UpdateAt:    result.UpdatedAt.Time,
		}
	}

	return tasks, nil
}
func (s *PGTaskStore) UpdateTask(ctx context.Context, id int32, title, description string, priority int32) (store.Task, error) {
	result, err := s.Queries.UpdateTask(ctx, UpdateTaskParams{
		ID:          id,
		Title:       title,
		Description: description,
		Priority:    priority,
	})
	if err != nil {
		return store.Task{}, err
	}

	return store.Task{
		Id:          result.ID,
		Title:       result.Title,
		Priority:    result.Priority,
		Description: result.Description,
		CreatedAt:   result.CreatedAt.Time,
		UpdateAt:    result.UpdatedAt.Time,
	}, nil
}

func (s *PGTaskStore) DeleteTask(ctx context.Context, id int32) error {
	return s.Queries.DeleteTask(ctx, id)
}
