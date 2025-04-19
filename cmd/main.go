package main

import (
	"fmt"

	"github.com/krijebr/todo-list/internal/entity"
)

func main() {
	task := entity.Task{
		Id:   123,
		Name: "test",
	}

	fmt.Printf("%+v", task)
}
