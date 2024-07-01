package data

import (
	"Mini-project/GoLang/interfaces"
	"Mini-project/GoLang/model"
	"context"
	"database/sql"
	"fmt"
)

type TaskDataLayerImpl struct {
	Dbconn *sql.DB
}

func (tdl *TaskDataLayerImpl) GetTaskByID(ctx context.Context, taskID string) *model.Task {
	var task model.Task
	query := "SELECT id, title, description, status FROM tasks WHERE id = ?"
	err := tdl.Dbconn.QueryRowContext(ctx, query, taskID).Scan(&task.ID, &task.Title, &task.Description, &task.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No task found with ID:", taskID)
			return nil
		}
		fmt.Println("Error querying database:", err)
		return nil
	}
	return &task
}

func (tdl *TaskDataLayerImpl) CreateTask(ctx context.Context, task *model.Task) error {
	query := "INSERT INTO tasks (title, description, status) VALUES (?, ?, ?)"
	_, err := tdl.Dbconn.ExecContext(ctx, query, task.Title, task.Description, "pending")
	return err
}

func NewTaskDataLayerImpl(conn *sql.DB) interfaces.TaskDataLayer {
	return &TaskDataLayerImpl{Dbconn: conn}
}
