package v1

import (
	"encoding/json"
	"io"
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
	tasks, err := h.usecase.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Ошибка получения ", err)
		return
	}
	data, err := json.Marshal(tasks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Ошибка кодирования данных в json ", err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	log.Println("Получение всех задач")
}

func (h *TaskHandlers) createTask(w http.ResponseWriter, r *http.Request) {
	var newtask entity.Task
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Ошибка чтения тела запроса ", err)
		return
	}
	err = json.Unmarshal(data, &newtask)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Ошибка декодирования тела запроса", err)
		return
	}
	newtask.IsDone = false
	err = h.usecase.Create(&newtask)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Ошибка создания ", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	log.Println("Задача создана")
}

func (h *TaskHandlers) deleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := parseId(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Ошибка обработки id ", err)
		return
	}
	err = h.usecase.DeleteById(id)
	if err != nil {
		switch {
		case err == usecase.ErrTaskNotFound:
			w.WriteHeader(http.StatusBadRequest)
			log.Println("Строки с таким id не существует", err)
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Ошибка удаления ", err)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	log.Println("Задача ", id, " удалена")
}

func (h *TaskHandlers) updateTask(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Name string `json:"name"`
	}
	id, err := parseId(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Ошибка обработки id ", err)
		return
	}
	req := request{}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Ошибка чтения тела запроса ", err)
		return
	}
	err = json.Unmarshal(data, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Ошибка декодирования тела запроса", err)
		return
	}
	err = h.usecase.UpdateNameById(id, req.Name)
	if err != nil {
		switch {
		case err == usecase.ErrTaskNotFound:
			w.WriteHeader(http.StatusBadRequest)
			log.Println("Строки с таким id не существует", err)
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Ошибка обновления ", err)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	log.Println("Задача ", id, " обновлена")
}

func (h *TaskHandlers) TaskSetDone(w http.ResponseWriter, r *http.Request) {
	id, err := parseId(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Ошибка обработки id ", err)
		return
	}
	err = h.usecase.SetDoneById(id)
	if err != nil {
		switch {
		case err == usecase.ErrTaskNotFound:
			w.WriteHeader(http.StatusBadRequest)
			log.Println("Строки с таким id не существует", err)
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Ошибка обновления ", err)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	log.Println("Задача ", id, " помечена сделанной")
}

func (h *TaskHandlers) TaskUnsetDone(w http.ResponseWriter, r *http.Request) {
	id, err := parseId(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Ошибка обработки id ", err)
		return
	}
	err = h.usecase.UnsetDoneById(id)
	if err != nil {
		switch {
		case err == usecase.ErrTaskNotFound:
			w.WriteHeader(http.StatusBadRequest)
			log.Println("Строки с таким id не существует", err)
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Ошибка обновления ", err)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	log.Println("Задача ", id, " помечена несделанной")
}
