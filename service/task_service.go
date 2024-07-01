package service

import (
	"Mini-project/GoLang/interfaces"
	"Mini-project/GoLang/model"
	"context"
	"fmt"
)

type TaskServiceImpl struct {
	TaskDataLayer interfaces.TaskDataLayer
}

func (ts *TaskServiceImpl) GetTaskByID(ctx context.Context, taskID string) *model.Task {
	fmt.Println("** Inside Task Service **")
	task := ts.TaskDataLayer.GetTaskByID(ctx, taskID)
	if task == nil {
		return nil
	}
	fmt.Printf("Task ID: %s, Task Title: %s, Task Description: %s, Task Status: %s\n", task.ID, task.Title, task.Description, task.Status)
	return task
}

func (ts *TaskServiceImpl) CreateTask(ctx context.Context, task *model.Task) error {
	return ts.TaskDataLayer.CreateTask(ctx, task)
}

func NewTaskServiceImpl(taskDL interfaces.TaskDataLayer) interfaces.TaskService {
	return &TaskServiceImpl{TaskDataLayer: taskDL}
}
