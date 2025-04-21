package v1

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/krijebr/todo-list/internal/entity"
	"github.com/krijebr/todo-list/internal/usecase"
)

type TaskHandlers struct {
	usecase usecase.TaskUseCase
}

func NewTaskHandlers(uc usecase.TaskUseCase) *TaskHandlers {
	return &TaskHandlers{usecase: uc}
}

func CreateRouter(uc usecase.TaskUseCase) *mux.Router {
	myhandler := NewTaskHandlers(uc)
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/task", myhandler.allTasks).Methods("GET")
	myRouter.HandleFunc("/task", myhandler.createTask).Methods("POST")
	myRouter.HandleFunc("/task/{id:[0-9]+}", myhandler.deleteTask).Methods("DELETE")
	myRouter.HandleFunc("/task/{id:[0-9]+}", myhandler.updateTask).Methods("PUT")
	myRouter.HandleFunc("/task/{id:[0-9]+}/set-done", myhandler.TaskSetDone).Methods("PUT")
	myRouter.HandleFunc("/task/{id:[0-9]+}/unset-done", myhandler.TaskUnsetDone).Methods("PUT")

	return myRouter
}

func (TH *TaskHandlers) allTasks(w http.ResponseWriter, r *http.Request) {
	_, err := TH.usecase.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Ошибка получения ", err)
	} else {
		w.WriteHeader(http.StatusOK)
		log.Println("Получение всех задач")
	}

}

func (TH *TaskHandlers) createTask(w http.ResponseWriter, r *http.Request) {
	t := new(entity.Task)
	err := TH.usecase.Create(t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Ошибка создания ", err)
	} else {
		w.WriteHeader(http.StatusCreated)
		log.Println("Задача создана")
	}
}

func (TH *TaskHandlers) deleteTask(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	err := TH.usecase.DeleteById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Ошибка удаления ", err)
	} else {
		w.WriteHeader(http.StatusOK)
		log.Println("Задача ", id, " удалена")
	}
}

func (TH *TaskHandlers) updateTask(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var newtaskname string
	err := TH.usecase.UpdateNameById(id, newtaskname)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Ошибка обновления ", err)
	} else {
		w.WriteHeader(http.StatusCreated)
		log.Println("Задача ", id, " обновлена")
	}
}

func (TH *TaskHandlers) TaskSetDone(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	err := TH.usecase.SetDoneById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Ошибка обновления ", err)
	} else {
		w.WriteHeader(http.StatusCreated)
		log.Println("Задача ", id, " помечена сделанной")
	}
}

func (TH *TaskHandlers) TaskUnsetDone(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	err := TH.usecase.UnsetDoneById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Ошибка обновления ", err)
	} else {
		w.WriteHeader(http.StatusCreated)
		log.Println("Задача ", id, " помечена несделанной")
	}
}
