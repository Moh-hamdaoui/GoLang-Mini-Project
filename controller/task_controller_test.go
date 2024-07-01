package controller

import (
	"Mini-project/GoLang/data"
	"Mini-project/GoLang/model"
	"Mini-project/GoLang/service"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetTask(t *testing.T) {

	mockTasks := []*model.Task{
		{ID: "2", Title: "Task 2", Description: "Description 2", Status: "pending"},
	}
	mockDataLayer := data.NewMockTaskDataLayer(mockTasks)
	taskService := service.NewTaskServiceImpl(mockDataLayer)
	taskController := &TaskController{taskService: taskService}

	req, err := http.NewRequest("GET", "/task/2", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/task/{id}", taskController.GetTask)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := model.Task{ID: "2", Title: "Task 2", Description: "Description 2", Status: "pending"}

	var actual model.Task
	err = json.NewDecoder(rr.Body).Decode(&actual)
	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestCreateTask(t *testing.T) {

	mockDataLayer := data.NewMockTaskDataLayer(nil)
	taskService := service.NewTaskServiceImpl(mockDataLayer)
	taskController := &TaskController{taskService: taskService}

	newTask := model.Task{ID: "2", Title: "New Task", Description: "New Description", Status: "pending"}
	taskJSON, _ := json.Marshal(newTask)

	req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer(taskJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/tasks", taskController.CreateTask)
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	createdTask := taskService.GetTaskByID(context.Background(), "2")
	if createdTask == nil || *createdTask != newTask {
		t.Errorf("expected %v, got %v", newTask, createdTask)
	}
}
