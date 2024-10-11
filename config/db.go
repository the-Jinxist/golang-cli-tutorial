package config

import (
	"log"
	"os"
	"path/filepath"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/jmoiron/sqlx"
	gap "github.com/muesli/go-app-paths"
)

var db *sqlx.DB

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
func initDB() {
	res, err := sqlx.Open("sqlite3", filepath.Join(setupPath(), "tasks.db"))
	if err != nil {
		log.Fatalf("error while opening db: %s", err)
	}

	setupDb(res)

	db = res
}

func setupDb(res *sqlx.DB) {
	if _, err := res.Query("SELECT * FROM tasks"); err != nil {
		if _, err := res.Exec(`CREATE TABLE "tasks" ( "id" INTEGER, "name" TEXT NOT NULL, "project" TEXT, "status" TEXT, "created_at" DATETIME, "updated_at" DATETIME, PRIMARY KEY("id" AUTOINCREMENT))`); err != nil {
			log.Fatalf("error while setting up db: %s", err)
		}
	}
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

func GetDB() *sqlx.DB {
	if db == nil {
		initDB()
	}

	return db
}
