package cmd

import (
	"database/sql"

	"github.com/the-Jinxist/golang-cli-tutorial/config"
)

type repository struct {
	db *sql.DB
}

func GetRepo() repository {
	return repository{
		db: config.GetDB(),
	}
}

func (r repository) getTasks() string {
	return "---------------------- --- --- --- \n(1) I wan chop\n(2) I wan sleep \n---------------------- --- --- ---"
}
