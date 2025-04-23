package repo

import (
	"database/sql"
	"log"

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
	_, err := r.db.Exec("insert into tasks (name) values ($1)", t.Name)
	if err != nil {
		log.Println("Ошибка добавления в базу данных", err)
		return err
	}
	/*c, err := result.RowsAffected()
	log.Printf("В базу данных добавлено %d строк", c)*/
	return nil
}
func (r *TaskRepoPg) GetAll() ([]*entity.Task, error) {
	rows, err := r.db.Query("select * from tasks order by id")
	if err != nil {
		return nil, err
	}
	tasks := []*entity.Task{}
	for rows.Next() {
		task := new(entity.Task)
		err := rows.Scan(&task.Id, &task.Name, &task.IsDone)
		if err != nil {
			log.Println("Ошибка чтения строки", err)
			continue
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
func (r *TaskRepoPg) DeleteById(id int) error {
	_, err := r.db.Exec("delete from tasks where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
func (r *TaskRepoPg) UpdateTaskById(id int, name string) error {
	_, err := r.db.Exec("update tasks set name = $1 where id = $2", name, id)
	if err != nil {
		return err
	}
	return nil
}
func (r *TaskRepoPg) SetDoneById(id int) error {
	_, err := r.db.Exec("update tasks set isdone = true where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
func (r *TaskRepoPg) UnsetDoneById(id int) error {
	_, err := r.db.Exec("update tasks set isdone = false where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
