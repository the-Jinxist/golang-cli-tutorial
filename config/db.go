package config

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	gap "github.com/muesli/go-app-paths"
)

var db *sql.DB

func setupPath() string {
	// get XDG paths
	scope := gap.NewScope(gap.User, "tasks")
	dirs, err := scope.DataDirs()
	if err != nil {
		log.Fatal(err)
	}
	// create the app base dir, if it doesn't exist
	var taskDir string
	if len(dirs) > 0 {
		taskDir = dirs[0]
	} else {
		taskDir, _ = os.UserHomeDir()
	}
	if err := initTaskDir(taskDir); err != nil {
		log.Fatal(err)
	}
	return taskDir
}

// openDB opens a SQLite database and stores that database in our special spot.
func InitDB() {
	res, err := sql.Open("sqlite3", filepath.Join(setupPath(), "tasks.db"))
	if err != nil {
		log.Fatalf("error while opening db")
	}

	db = res
}

func initTaskDir(path string) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return os.Mkdir(path, 0o770)
		}
		return err
	}
	return nil
}

func GetDB() *sql.DB {
	return db
}
