package repo

import (
	"database/sql"

	"github.com/krijebr/todo-list/internal/entity"
	_ "github.com/lib/pq"
)

type TaskRepoPg struct {
	db *sql.DB
}

func NewTaskRepoPg(db *sql.DB) TaskRepository {
	return &TaskRepoPg{
		db: db}
}

func (r *TaskRepoPg) Create(t *entity.Task) error {
	return nil
}
func (r *TaskRepoPg) GetAll() ([]*entity.Task, error) {
	return nil, nil
}
func (r *TaskRepoPg) DeleteById(id int) error {
	return nil
}
func (r *TaskRepoPg) UpdateTaskById(id int, name string) error {
	return nil
}
func (r *TaskRepoPg) SetDoneById(id int) error {
	return nil
}
func (r *TaskRepoPg) UnsetDoneById(id int) error {
	return nil
}
