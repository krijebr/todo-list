package usecase

import (
	"errors"

	"github.com/krijebr/todo-list/internal/entity"
	"github.com/krijebr/todo-list/internal/repo"
)

type Task struct {
	repo repo.TaskRepository
}

func NewTaskUseCase(r repo.TaskRepository) TaskUseCase {
	return &Task{repo: r}
}

var err = errors.New("not implemented")

func (t *Task) Create(task *entity.Task) error {
	err := t.repo.Create(task)
	if err != nil {
		return err
	}
	return nil
}
func (t *Task) GetAll() ([]*entity.Task, error) {
	var tasks []*entity.Task
	tasks, err := t.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
func (t *Task) DeleteById(id int) error {
	_, err := t.repo.GetById(id)
	if err != nil {
		switch {
		case err == repo.ErrTaskNotFound:
			return ErrTaskNotFound
		default:
			return err
		}
	}
	err = t.repo.DeleteById(id)
	if err != nil {
		return err
	}
	return nil
}
func (t *Task) SetDoneById(id int) error {
	_, err := t.repo.GetById(id)
	if err != nil {
		switch {
		case err == repo.ErrTaskNotFound:
			return ErrTaskNotFound
		default:
			return err
		}
	}
	err = t.repo.SetDoneById(id)
	if err != nil {
		return err
	}
	return nil
}
func (t *Task) UnsetDoneById(id int) error {
	_, err := t.repo.GetById(id)
	if err != nil {
		switch {
		case err == repo.ErrTaskNotFound:
			return ErrTaskNotFound
		default:
			return err
		}
	}
	err = t.repo.UnsetDoneById(id)
	if err != nil {
		return err
	}
	return nil
}
func (t *Task) UpdateNameById(id int, name string) error {
	_, err := t.repo.GetById(id)
	if err != nil {
		switch {
		case err == repo.ErrTaskNotFound:
			return ErrTaskNotFound
		default:
			return err
		}
	}
	err = t.repo.UpdateTaskById(id, name)
	if err != nil {
		return err
	}
	return nil
}
