package services

import (
	"log/slog"
	"todolist/backend/internal/repository"
	"todolist/backend/models"
)

type TaskService struct {
	taskRepo *repository.TaskRepository
	logger   *slog.Logger
}

func NewTaskService(repo *repository.TaskRepository, logger *slog.Logger) *TaskService {
	return &TaskService{taskRepo: repo, logger: logger}
}

func (s *TaskService) GetTasksService() ([]models.Task, error) {
	return s.taskRepo.GetTasks()
}

func (s *TaskService) InsertTaskService(title, description, deadline, priority string) error {
	return s.taskRepo.InsertTask(title, description, deadline, priority)
}

func (s *TaskService) UpdateTaskService(id int, title, description, deadline, priority string) error {
	return s.taskRepo.UpdateTask(id, title, description, deadline, priority)
}

func (s *TaskService) DeleteTaskService(id int) error {
	return s.taskRepo.DeleteTask(id)
}
