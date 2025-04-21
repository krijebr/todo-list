package v1

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/task", allTasks).Methods("GET")
	myRouter.HandleFunc("/task", createTask).Methods("POST")
	myRouter.HandleFunc("/task/{id:[0-9]+}", deleteTask).Methods("DELETE")
	myRouter.HandleFunc("/task/{id:[0-9]+}", updateTask).Methods("PUT")

	return myRouter
}

func allTasks(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Получение всех задач"))
	log.Println("Получение всех задач")
}

func createTask(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Задача создана"))
	log.Println("Задача создана")
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Задача " + mux.Vars(r)["id"] + " удалена"))
	log.Println("Задача " + mux.Vars(r)["id"] + " удалена")
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Задача " + mux.Vars(r)["id"] + " обновлена"))
	log.Println("Задача " + mux.Vars(r)["id"] + " обновлена")

}
