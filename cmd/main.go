package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	v1 "github.com/krijebr/todo-list/internal/controller/http/v1"
	"github.com/krijebr/todo-list/internal/repo"
	"github.com/krijebr/todo-list/internal/usecase"
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

	connectionDbUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, portdb, username, password, dbname)
	rep1, err := repo.NewTaskRepoInDb(connectionDbUrl)
	if err != nil {
		log.Println(err)
	}
	_ = rep1
	rep := repo.NewTaskRepoInMemory()
	uc := usecase.NewTaskUseCase(rep)
	r := v1.CreateRouter(uc)

	adr := ":" + strconv.Itoa(port)
	err = http.ListenAndServe(adr, r)
	if err != nil {
		log.Println("Ошибка запуска сервера", err)
	}
}
