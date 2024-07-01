package main

import (
	"database/sql"
	"log"
	"Mini-project/GoLang/controller"
	"Mini-project/GoLang/data"
	"Mini-project/GoLang/service"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		description TEXT,
		status TEXT
	);`)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	initDB()

	r := mux.NewRouter()

	taskDataLayer := data.NewTaskDataLayerImpl(db)

	taskService := service.NewTaskServiceImpl(taskDataLayer)

	controller.NewTaskController(r, taskService)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to the Task Manager API"))
	}).Methods("GET")

	log.Println("Listening on :8080...")
	http.ListenAndServe(":8080", r)
}
