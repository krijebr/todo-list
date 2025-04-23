package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	v1 "github.com/krijebr/todo-list/internal/controller/http/v1"
	"github.com/krijebr/todo-list/internal/repo"
	"github.com/krijebr/todo-list/internal/usecase"
	_ "github.com/lib/pq"
)

const port int = 8080

const (
	host     = "localhost"
	portdb   = 5432
	username = "postgres"
	password = "mysecretpassword"
	dbname   = "postgres"
)

func main() {
	log.Println("starting app")

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, portdb, username, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Ошибка подключения к базе данных", err)
		return
	}
	_, err = db.Exec("select 1")
	if err != nil {
		log.Println("Ошибка работы с базой данных", err)
		return
	}
	rep1 := repo.NewTaskRepoPg(db)
	_ = rep1
	rep := repo.NewTaskRepoInMemory()
	_ = rep
	uc := usecase.NewTaskUseCase(rep1)
	r := v1.CreateRouter(uc)

	adr := ":" + strconv.Itoa(port)
	err = http.ListenAndServe(adr, r)
	if err != nil {
		log.Println("Ошибка запуска сервера", err)
	}
}
