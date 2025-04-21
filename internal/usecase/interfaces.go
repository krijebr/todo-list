package usecase

import "github.com/krijebr/todo-list/internal/entity"

type TaskUseCase interface {
	Create(task *entity.Task) error
	GetAll() ([]*entity.Task, error)
	DeleteById(id int) error
	SetDoneById(id int) error
	UnsetDoneById(id int) error
	UpdateNameById(id int, name string) error
}
