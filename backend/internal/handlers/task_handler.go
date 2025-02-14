package handlers

import (
	"log/slog"
	myerrors "todolist/backend/internal/errors"
	"todolist/backend/internal/services"
	"todolist/backend/models"
)

// handle task-related operations
type TaskHandler struct {
	taskService *services.TaskService
	logger      *slog.Logger
}

func NewTaskHandler(service *services.TaskService, logger *slog.Logger) *TaskHandler {
	return &TaskHandler{taskService: service, logger: &slog.Logger{}}
}

func (h *TaskHandler) GetTasksHandler() ([]models.Task, error) {
	tasks, err := h.taskService.GetTasksService()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (h *TaskHandler) InsertTasksHandler(title, description, deadline, priority, status string) error {
	if title == "" {
		h.logger.Error("Failed to insert task", "error", myerrors.ErrEmptyTitle)
		return myerrors.ErrEmptyTitle
	}

	return h.taskService.InsertTaskService(title, description, deadline, priority, status)
}

func (h *TaskHandler) UpdateTasksHandler(id int, priority, status string) error {
	return h.taskService.UpdateTaskService(id, priority, status)
}

func (h *TaskHandler) DeleteTasksHandler(id int) error {
	return h.taskService.DeleteTaskService(id)
}
