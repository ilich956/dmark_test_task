package repository

import (
	"database/sql"
	"log/slog"
	"os"
	"todolist/backend/models"
)

type TaskRepository struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewTaskRepository(db *sql.DB, logger *slog.Logger) *TaskRepository {
	return &TaskRepository{db: db, logger: logger}
}

func (r *TaskRepository) InitDB() error {
	query := `CREATE TABLE IF NOT EXISTS tasks(
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    title TEXT NOT NULL,
	    description TEXT,
		deadline    DATETIME NOT NULL,
		priority    TEXT NOT NULL CHECK(priority IN ('Low', 'Medium', 'High')),
		status      TEXT NOT NULL CHECK(status IN ('Pending', 'In progress', 'Completed')) DEFAULT 'Pending'
	);`
	_, err := r.db.Exec(query)
	if err != nil {
		r.logger.Error("Failed to create table", "error", err)
		os.Exit(1)
	}
	return nil
}

func (r *TaskRepository) GetTasks() ([]models.Task, error) {
	rows, err := r.db.Query("SELECT * FROM tasks")
	if err != nil {
		r.logger.Error("Failed to execute SELECT query", "error", err)
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task

		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Deadline, &task.Priority, &task.Status)
		if err != nil {
			r.logger.Error("Failed to scan rows", "error", err)
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *TaskRepository) InsertTask(title, description, deadline, priority, status string) error {
	_, err := r.db.Exec("INSERT INTO tasks (title, description, deadline, priority, status) VALUES (?,?,?,?,?)", title, description, deadline, priority, status)
	if err != nil {
		r.logger.Error("Failed to execute INSERT query", "error", err)
		return err
	}
	return nil
}

func (r *TaskRepository) UpdateTask(id int, title, description, deadline, priority string) error {
	_, err := r.db.Exec("UPDATE tasks SET title = ?, description = ?, deadline = ?, priority = ? WHERE id = ?", title, description, deadline, priority, id)
	if err != nil {
		r.logger.Error("Failed to execute UPDATE query", "error", err)
		return err
	}
	return nil
}

func (r *TaskRepository) DeleteTask(id int) error {
	if _, err := r.db.Exec("DELETE FROM tasks WHERE id = ?", id); err != nil {
		r.logger.Error("Failed to execute DELETE query", "error", err)
		return err
	}
	return nil
}
