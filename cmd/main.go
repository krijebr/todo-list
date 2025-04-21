package main

import (
	"log"
	"net/http"
	"strconv"

	v1 "github.com/krijebr/todo-list/internal/controller/http/v1"
)

const port int = 8080

func main() {

	r := v1.CreateRouter()

	adr := ":" + strconv.Itoa(port)
	err := http.ListenAndServe(adr, r)
	if err != nil {
		log.Println("Ошибка запуска сервера", err)
	}
}
