package backend

import (
	"context"
	"database/sql"
	"log/slog"
	"todolist/backend/internal/config"
	"todolist/backend/internal/handlers"
	"todolist/backend/internal/repository"
	"todolist/backend/internal/services"
	"todolist/backend/models"
)

// App struct
type App struct {
	taskHandler *handlers.TaskHandler
	db          *sql.DB
	logger      *slog.Logger
}

// NewApp creates a new App application struct
func NewApp() *App {
	logger := config.NewLogger()

	return &App{logger: logger}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	db, err := sql.Open("sqlite3", "backend/database/tasks.db")
	if err != nil {
		a.logger.Error("Failed to open database", "error", err)
	}
	a.db = db

	repo := repository.NewTaskRepository(a.db, a.logger)
	repo.InitDB()
	service := services.NewTaskService(repo, a.logger)
	handler := handlers.NewTaskHandler(service, a.logger)

	a.taskHandler = handler
}

func (a *App) Shutdown(ctx context.Context) {
	if a.db != nil {
		if err := a.db.Close(); err != nil {
			a.logger.Error("Failed to close database", "error", err)
		} else {
			a.logger.Info("Database closed succesfully")
		}
	}
}

func (a *App) GetTasks() ([]models.Task, error) {
	return a.taskHandler.GetTasksHandler()
}

func (a *App) InsertTask(title, description, deadline, priority string) error {
	return a.taskHandler.InsertTasksHandler(title, description, deadline, priority)
}
