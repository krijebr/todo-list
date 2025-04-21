package usecase

import (
	"errors"

	"github.com/krijebr/todo-list/internal/entity"
)

type Task struct {
}

var err = errors.New("not implemented")

func (t *Task) Create(task *entity.Task) error {
	return err
}
func (t *Task) GetAll() ([]*entity.Task, error) {
	return nil, err
}
func (t *Task) DeleteById(id int) error {
	return err
}
func (t *Task) SetDoneById(id int) error {
	return err
}
func (t *Task) UnsetDoneById(id int) error {
	return err
}
func (t *Task) UpdateNameById(id int, name string) error {
	return err
}

func NewTaskUseCase() TaskUseCase {
	t := new(Task)
	return t
}
