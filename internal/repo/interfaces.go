package repo

import "github.com/krijebr/todo-list/internal/entity"

type TaskRepository interface {
	Create(t *entity.Task) error
	GetAll() ([]*entity.Task, error)
	DeleteById(id int) error
	UpdateTaskById(id int, name string) error
	SetDoneById(id int) error
	UnsetDoneById(id int) error
}
