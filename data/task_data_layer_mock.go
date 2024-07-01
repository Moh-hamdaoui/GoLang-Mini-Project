package data

import (
	"Mini-project/GoLang/interfaces"
	"Mini-project/GoLang/model"
	"context"
)

type MockTaskDataLayer struct {
	Tasks []*model.Task
}

func (m *MockTaskDataLayer) GetTaskByID(ctx context.Context, taskID string) *model.Task {
	for _, task := range m.Tasks {
		if task.ID == taskID {
			return task
		}
	}
	return nil
}

func (m *MockTaskDataLayer) CreateTask(ctx context.Context, task *model.Task) error {
	m.Tasks = append(m.Tasks, task)
	return nil
}

func NewMockTaskDataLayer(tasks []*model.Task) interfaces.TaskDataLayer {
	return &MockTaskDataLayer{Tasks: tasks}
}
