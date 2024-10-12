package cmd

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func GetRepo(db *sqlx.DB) repository {
	return repository{
		db: db,
	}
}

func (t *repository) getTasks(project, _ string) ([]Task, error) {
	var tasks []Task

	//TODO: Need to add sql query builder to add status to query dynamically
	query := `SELECT * FROM tasks`
	if len(project) > 3 {
		query += fmt.Sprintf(` WHERE project = "%s"`, project)
	}
	err := t.db.Select(&tasks, query)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("you have not created any tasks yet. Use `add --name` to create a new task")
	}

	return tasks, err
}

func (t *repository) finishTask(id int) (string, error) {
	var name string
	err := t.db.QueryRow(`update tasks set status = $1, updated_at = $2 where id = $3 returning name`, "finished", time.Now(), id).Scan(&name)
	if err == sql.ErrNoRows {
		return name, fmt.Errorf("no tasks with id: %d", id)
	}
	return name, err
}

func (t *repository) startTask(id int) (string, error) {
	var name string
	err := t.db.QueryRow(`update tasks set status = $1, updated_at = $2 where id = $3 returning name`, "in_progress", time.Now(), id).Scan(&name)
	if err == sql.ErrNoRows {
		return name, fmt.Errorf("no tasks with id: %d", id)
	}
	return name, err
}

func (t *repository) deleteTask(id int) (string, error) {
	var name string
	err := t.db.QueryRow(`delete from tasks where id = $1 returning name`, id).Scan(&name)
	if err == sql.ErrNoRows {
		return name, fmt.Errorf("no tasks with id: %d", id)
	}
	return name, err
}

func (t *repository) clearAllTasks() error {

	_, err := t.db.Exec(`delete from tasks`)
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}

func (t *repository) createTask(task Task) error {

	if task.Project == "" {
		task.Project = "default"
	}

	now := time.Now()
	_, err := t.db.Exec(`insert into tasks (name, project, status, created_at, updated_at)
		values ($1, $2, $3, $4, $5)`, task.Name, task.Project, task.Status, now, now)
	return err
}
