package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/krijebr/todo-list/internal/config"
	v1 "github.com/krijebr/todo-list/internal/controller/http/v1"
	"github.com/krijebr/todo-list/internal/repo"
	"github.com/krijebr/todo-list/internal/usecase"
	_ "github.com/lib/pq"
)

const confpath string = "./config/config.json"

func main() {
	log.Println("starting app")

	cfg, err := config.InitConfigFromJson(confpath)
	if err != nil {
		log.Println("Ошибка инициализации", err)
		return
	}

	db, err := initDB(&cfg.Postgres)
	if err != nil {
		log.Println("Ошибка инициализации базы данных", err)
		return
	}
	repoPostgres := repo.NewTaskRepoPg(db)
	_ = repo.NewTaskRepoInMemory()
	uc := usecase.NewTaskUseCase(repoPostgres)
	r := v1.CreateRouter(uc)

	addr := ":" + strconv.Itoa(cfg.HttpServer.Port)
	log.Printf("starting http server on port %s\n", addr)
	err = http.ListenAndServe(addr, r)
	if err != nil {
		log.Println("Ошибка запуска сервера", err)
	}
}

func initDB(cfg *config.Postgres) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.UserName, cfg.Password, cfg.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec("select 1")
	if err != nil {
		return nil, err
	}
	return db, nil
}
