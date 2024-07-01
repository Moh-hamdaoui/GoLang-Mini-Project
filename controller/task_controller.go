package controller

import (
	"Mini-project/GoLang/interfaces"
	"Mini-project/GoLang/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type TaskController struct {
	taskService interfaces.TaskService
}

func (tc *TaskController) GetTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("** Inside Task Controller **")
	requestCtx := r.Context()
	vars := mux.Vars(r)
	taskID := vars["id"]
	fmt.Println("Task ID:", taskID)
	task := tc.taskService.GetTaskByID(requestCtx, taskID)
	if task == nil {
		fmt.Println("Task not found")
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func (tc *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = tc.taskService.CreateTask(r.Context(), &task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func NewTaskController(r *mux.Router, taskServiceObj interfaces.TaskService) {
	taskController := &TaskController{
		taskService: taskServiceObj,
	}

	r.HandleFunc("/task/{id}", taskController.GetTask).Methods("GET")
	r.HandleFunc("/tasks", taskController.CreateTask).Methods("POST")
}
