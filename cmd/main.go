package main

import (
	"log"
	"net/http"
	"strconv"

	v1 "github.com/krijebr/todo-list/internal/controller/http/v1"
	"github.com/krijebr/todo-list/internal/usecase"
)

const port int = 8080

func main() {

	uc := usecase.NewTaskUseCase()
	r := v1.CreateRouter(uc)

	adr := ":" + strconv.Itoa(port)
	err := http.ListenAndServe(adr, r)
	if err != nil {
		log.Println("Ошибка запуска сервера", err)
	}
}
