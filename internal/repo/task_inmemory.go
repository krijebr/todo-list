package repo

import (
	"github.com/krijebr/todo-list/internal/entity"
)

type TaskRepoInMemory struct {
	tasks map[int]*entity.Task
	id    int
}

func NewTaskRepoInMemory() TaskRepository {
	return &TaskRepoInMemory{
		tasks: make(map[int]*entity.Task),
		id:    1}
}

func (r *TaskRepoInMemory) newId() int {
	b := true
	for b {
		if _, inMap := r.tasks[r.id]; inMap {
			r.id++
		} else {
			b = false
		}
	}
	return r.id
}
func (r *TaskRepoInMemory) Create(t *entity.Task) error {
	id := r.newId()
	t.Id = id
	r.tasks[id] = t
	return nil
}
func (r *TaskRepoInMemory) GetAll() ([]*entity.Task, error) {
	var tasks []*entity.Task
	for _, value := range r.tasks {
		tasks = append(tasks, value)
	}
	return tasks, nil
}
func (r *TaskRepoInMemory) DeleteById(id int) error {
	delete(r.tasks, id)
	return nil
}
func (r *TaskRepoInMemory) UpdateTaskById(id int, name string) error {
	r.tasks[id].Name = name
	return nil
}
func (r *TaskRepoInMemory) SetDoneById(id int) error {
	r.tasks[id].IsDone = true
	return nil
}
func (r *TaskRepoInMemory) UnsetDoneById(id int) error {
	r.tasks[id].IsDone = false
	return nil
}

func (r *TaskRepoInMemory) GetById(id int) (*entity.Task, error) {
	if task, inMap := r.tasks[id]; inMap {
		return task, nil
	} else {
		return nil, ErrTaskNotFound
	}
}
