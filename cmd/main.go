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

type ConfigPostgres struct {
	host     string
	portdb   int
	username string
	password string
	dbname   string
}

func main() {
	log.Println("starting app")

	configpostgres := ConfigPostgres{
		host:     host,
		portdb:   portdb,
		username: username,
		password: password,
		dbname:   dbname,
	}
	db, err := initDB(configpostgres)
	if err != nil {
		log.Println("Ошибка инициализации базы данных", err)
		return
	}
	repoPostgres := repo.NewTaskRepoPg(db)
	_ = repo.NewTaskRepoInMemory()
	uc := usecase.NewTaskUseCase(repoPostgres)
	r := v1.CreateRouter(uc)

	addr := ":" + strconv.Itoa(port)
	log.Printf("starting http server on port %s\n", addr)
	err = http.ListenAndServe(addr, r)
	if err != nil {
		log.Println("Ошибка запуска сервера", err)
	}
}

func initDB(cp ConfigPostgres) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cp.host, cp.portdb, cp.username, cp.password, cp.dbname)

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
