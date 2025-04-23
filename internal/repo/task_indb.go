package repo

import (
	"database/sql"

	"github.com/krijebr/todo-list/internal/entity"
	_ "github.com/lib/pq"
)

type TaskRepoInDb struct {
	db *sql.DB
}

func NewTaskRepoInDb(connectionUrl string) (TaskRepository, error) {
	var r TaskRepoInDb
	var err error
	r.db, err = sql.Open("postgres", connectionUrl)
	if err != nil {
		return nil, err
	}
	defer r.db.Close()
	err = r.db.Ping()
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (r *TaskRepoInDb) Create(t *entity.Task) error {
	return nil
}
func (r *TaskRepoInDb) GetAll() ([]*entity.Task, error) {
	return nil, nil
}
func (r *TaskRepoInDb) DeleteById(id int) error {
	return nil
}
func (r *TaskRepoInDb) UpdateTaskById(id int, name string) error {
	return nil
}
func (r *TaskRepoInDb) SetDoneById(id int) error {
	return nil
}
func (r *TaskRepoInDb) UnsetDoneById(id int) error {
	return nil
}
