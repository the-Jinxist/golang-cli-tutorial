package cmd

import (
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

func (t *repository) getTasks(project string) ([]Task, error) {
	var tasks []Task

	query := `SELECT * FROM tasks`
	if len(project) > 3 {
		query += fmt.Sprintf(` WHERE project = "%s"`, project)
	}
	err := t.db.Select(&tasks, query)
	return tasks, err
}

func (t *repository) finishTask(id int) error {
	_, err := t.db.Exec(`update tasks set status = $1 where id = $2`, "finished", id)
	return err
}

func (t *repository) deleteTask(id int) error {
	_, err := t.db.Exec(`delete tasks where id = $2`, "finished", id)
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
