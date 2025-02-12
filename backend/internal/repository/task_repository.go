package repository

import (
	"database/sql"
	"log/slog"
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
	priority    TEXT NOT NULL CHECK(status IN ('Low', 'Medium', 'High'),
	status      TEXT NOT NULL CHECK(status IN ('Pending', 'In progress', 'Completed'))
)`
	_, err := r.db.Exec(query)
	return err
}

func (r *TaskRepository) GetTasks() ([]models.Task, error) {
	rows, err := r.db.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task

		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Deadline, &task.Priority, &task.Status)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *TaskRepository) InsertTask(title, description, deadline, priority string) error {
	_, err := r.db.Exec("INSERT INTO tasks (title, description, deadline, priority) VALUES (?,?,?,?)", title, description, deadline, priority)
	return err
}

func (r *TaskRepository) UpdateTask(id int, title, description, deadline, priority string) error {
	_, err := r.db.Exec("UPDATE tasks SET title = ?, description = ?, deadline = ?, priority = ? WHERE id = ?", title, description, deadline, priority, id)
	return err
}

func (r *TaskRepository) DeleteTask(id int) error {
	if _, err := r.db.Exec("DELETE FROM tasks WHERE id = ?", id); err != nil {
		return err
	}
	return nil
}
