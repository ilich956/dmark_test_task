package repository

import (
	"database/sql"
	"todolist/backend/models"
)

type TaskRepository struct {
	DB *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (r *TaskRepository) GetTasks() ([]models.Task, error) {
	rows, err := r.DB.Query("SELECT * FROM tasks")
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
	_, err := r.DB.Exec("INSERT INTO tasks (title, description, deadline, priority) VALUES (?,?,?,?)", title, description, deadline, priority)
	return err
}

func (r *TaskRepository) UpdateTask(id int, title, description, deadline, priority string) error {
	_, err := r.DB.Exec("UPDATE tasks SET title = ?, description = ?, deadline = ?, priority = ? WHERE id = ?", title, description, deadline, priority, id)
	return err
}

func (r *TaskRepository) DeleteTask(id int) error {
	if _, err := r.DB.Exec("DELETE FROM tasks WHERE id = ?", id); err != nil {
		return err
	}
}
