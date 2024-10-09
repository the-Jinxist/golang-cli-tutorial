package cmd

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/the-Jinxist/golang-cli-tutorial/config"
)

type repository struct {
	db *sqlx.DB
}

func GetRepo() repository {
	return repository{
		db: config.GetDB(),
	}
}

func (t *repository) getTasks(project string) ([]Task, error) {
	var tasks []Task

	query := `SELECT * FROM tasks`
	if len(project) > 3 {
		query += fmt.Sprintf(`WHERE project = %s`, project)
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

	if task.Project != "" {
		task.Project = "default"
	}

	_, err := t.db.Exec(`insert into tasks (name, project, status, created_at, updated_at)
		values ($1, $2, $3, now(), now())`, task.Name, task.Project, task.Status)
	return err
}
