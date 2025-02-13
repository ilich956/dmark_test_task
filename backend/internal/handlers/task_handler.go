package handlers

import (
	"log/slog"
	myerrors "todolist/backend/internal/errors"
	"todolist/backend/internal/services"
	"todolist/backend/models"
)

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
	//fmt.Println(tasks)
}

func (h *TaskHandler) InsertTasksHandler(title, description, deadline, priority, status string) error {
	if title == "" {
		h.logger.Error("Failed to insert task", "error", myerrors.ErrEmptyTitle)
		return myerrors.ErrEmptyTitle
	}

	//handle error

	return h.taskService.InsertTaskService(title, description, deadline, priority, status)
}

func (h *TaskHandler) UpdateTasksHandler(id int, title, description, deadline, priority string) error {
	if title == "" {
		h.logger.Error("Failed to insert task", "error", myerrors.ErrEmptyTitle)
		return myerrors.ErrEmptyTitle
	}

	//handle error
	return h.taskService.UpdateTaskService(id, title, description, deadline, priority)
}

func (h *TaskHandler) DeleteTasksHandler(id int) error {
	//handle error

	return h.taskService.DeleteTaskService(id)
}
