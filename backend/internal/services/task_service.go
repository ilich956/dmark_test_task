package services

import (
	"log/slog"
	"todolist/backend/internal/repository"
	"todolist/backend/models"
)

// Methods to task related operations
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

func (s *TaskService) InsertTaskService(title, description, deadline, priority, status string) error {
	return s.taskRepo.InsertTask(title, description, deadline, priority, status)
}

// REWRITE THIS, IT SHOULD UPDATE ONLY STATUS AND PRIORITY
func (s *TaskService) UpdateTaskService(id int, priority, status string) error {
	return s.taskRepo.UpdateTask(id, priority, status)
}

func (s *TaskService) DeleteTaskService(id int) error {
	return s.taskRepo.DeleteTask(id)
}
