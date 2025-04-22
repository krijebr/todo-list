package repo

import "github.com/krijebr/todo-list/internal/entity"

type TaskRepoInMemory struct {
	tasks map[int]*entity.Task
	Id    int
}

func NewId(T *TaskRepoInMemory) int {
	b := true
	for b {
		if _, inMap := T.tasks[T.Id]; inMap {
			T.Id++
		} else {
			b = false
		}
	}
	return T.Id
}
func (T *TaskRepoInMemory) Create(t *entity.Task) error {
	id := NewId(T)
	t.Id = id
	T.tasks[id] = t
	return nil
}
func (T *TaskRepoInMemory) GetAll() ([]*entity.Task, error) {
	var tasks []*entity.Task
	for _, value := range T.tasks {
		tasks = append(tasks, value)
	}
	return tasks, nil
}
func (T *TaskRepoInMemory) DeleteById(id int) error {
	delete(T.tasks, id)
	return nil
}
func (T *TaskRepoInMemory) UpdateTaskById(id int, name string) error {
	T.tasks[id].Name = name
	return nil
}
func (T *TaskRepoInMemory) SetDoneById(id int) error {
	T.tasks[id].IsDone = true
	return nil
}
func (T *TaskRepoInMemory) UnsetDoneById(id int) error {
	T.tasks[id].IsDone = false
	return nil
}
