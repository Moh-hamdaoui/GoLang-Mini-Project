package service

import (
	"Mini-project/GoLang/data"
	"Mini-project/GoLang/model"
	"context"
	"testing"
)

func TestGetTaskByID(t *testing.T) {
	mockTasks := []*model.Task{
		{ID: "1", Title: "Task 1", Description: "Description 1", Status: "pending"},
		{ID: "2", Title: "Task 2", Description: "Description 2", Status: "pending"},
	}
	mockDataLayer := data.NewMockTaskDataLayer(mockTasks)
	taskService := NewTaskServiceImpl(mockDataLayer)

	t.Run("existing task", func(t *testing.T) {
		task := taskService.GetTaskByID(context.Background(), "1")
		if task == nil || task.ID != "1" {
			t.Errorf("expected task ID 1, got %v", task)
		}
	})

	t.Run("non-existing task", func(t *testing.T) {
		task := taskService.GetTaskByID(context.Background(), "3")
		if task != nil {
			t.Errorf("expected nil, got %v", task)
		}
	})
}

func TestCreateTask(t *testing.T) {
	mockDataLayer := data.NewMockTaskDataLayer(nil)
	taskService := NewTaskServiceImpl(mockDataLayer)

	newTask := &model.Task{ID: "1", Title: "New Task", Description: "New Description", Status: "pending"}
	err := taskService.CreateTask(context.Background(), newTask)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	task := taskService.GetTaskByID(context.Background(), "1")
	if task == nil || task.Title != "New Task" {
		t.Errorf("expected 'New Task', got %v", task)
	}
}
