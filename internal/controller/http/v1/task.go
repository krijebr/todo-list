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
func parseId(r *http.Request) (int, error) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (h *TaskHandlers) allTasks(w http.ResponseWriter, r *http.Request) {
	_, err := h.usecase.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Ошибка получения ", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Println("Получение всех задач")

}

func (h *TaskHandlers) createTask(w http.ResponseWriter, r *http.Request) {
	t := new(entity.Task)
	err := h.usecase.Create(t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Ошибка создания ", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	log.Println("Задача создана")

}

func (h *TaskHandlers) deleteTask(w http.ResponseWriter, r *http.Request) {
	id, err1 := parseId(r)
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Ошибка обработки id ", err1)
		return
	}
	err := h.usecase.DeleteById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Ошибка удаления ", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Println("Задача ", id, " удалена")

}

func (h *TaskHandlers) updateTask(w http.ResponseWriter, r *http.Request) {
	id, err1 := parseId(r)
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Ошибка обработки id ", err1)
		return
	}
	var newtaskname string
	err := h.usecase.UpdateNameById(id, newtaskname)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Ошибка обновления ", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	log.Println("Задача ", id, " обновлена")

}

func (h *TaskHandlers) TaskSetDone(w http.ResponseWriter, r *http.Request) {
	id, err1 := parseId(r)
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Ошибка обработки id ", err1)
		return
	}
	err := h.usecase.SetDoneById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Ошибка обновления ", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	log.Println("Задача ", id, " помечена сделанной")

}

func (h *TaskHandlers) TaskUnsetDone(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	err := h.usecase.UnsetDoneById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Ошибка обновления ", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	log.Println("Задача ", id, " помечена несделанной")

}
