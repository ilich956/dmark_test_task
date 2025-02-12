package services

import (
	myerrors "todolist/backend/internal/errors"
	"todolist/backend/internal/repository"
	"todolist/backend/models"
)

type TaskService struct {
	taskRepo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{taskRepo: repo}
}

func (s *TaskService) GetTasksService() ([]models.Task, error) {
	return s.taskRepo.GetTasks()
}

func (s *TaskService) InsertTaskService(title, description, deadline, priority string) error {
	if title == "" {
		return myerrors.ErrEmptyTitle
	}

	return s.taskRepo.InsertTask(title, description, deadline, priority)
}

func (s *TaskService) UpdateTaskService(id int, title, description, deadline, priority string) error {
	if title == "" {
		return myerrors.ErrEmptyTitle
	}

	return s.taskRepo.UpdateTask(id, title, description, deadline, priority)
}

func (s *TaskService) DeleteTaskService(id int) error {
	return s.taskRepo.DeleteTask(id)
}
