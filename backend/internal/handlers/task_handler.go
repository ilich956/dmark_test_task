package handlers

import (
	myerrors "todolist/backend/internal/errors"
	"todolist/backend/internal/services"
	"todolist/backend/models"
)

type TaskHandler struct {
	taskService *services.TaskService
}

func NewTaskHandler(service *services.TaskService) *TaskHandler {
	return &TaskHandler{taskService: service}
}

func (h *TaskHandler) GetTasksHandler() ([]models.Task, error) {
	return h.taskService.GetTasksService()
}

func (h *TaskHandler) InsertTasksHandler(title, description, deadline, priority string) error {
	if title == "" {
		return myerrors.ErrEmptyTitle
	}

	return h.taskService.InsertTaskService(title, description, deadline, priority)
}

func (h *TaskHandler) UpdateTasksHandler(id int, title, description, deadline, priority string) error {
	if title == "" {
		return myerrors.ErrEmptyTitle
	}

	return h.taskService.UpdateTaskService(id, title, description, deadline, priority)
}

func (h *TaskHandler) DeleteTasksHandler(id int) error {
	return h.taskService.DeleteTaskService(id)
}
