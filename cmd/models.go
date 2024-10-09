package cmd

import "time"

type Task struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Project   string    `json:"project" db:"project"`
	Status    string    `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
