package interfaces

import (
	"Mini-project/GoLang/model"
	"context"
)

type TaskService interface {
	GetTaskByID(ctx context.Context, taskID string) *model.Task
	CreateTask(ctx context.Context, task *model.Task) error
}

type TaskDataLayer interface {
	GetTaskByID(ctx context.Context, taskID string) *model.Task
	CreateTask(ctx context.Context, task *model.Task) error
}
