package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var taskCommand = &cobra.Command{
	Use:   "tasks",
	Short: "This is a list of all your tasks",
	Run: func(cmd *cobra.Command, args []string) {

		repo := GetRepo()
		tasks := repo.getTasks()

		fmt.Print(tasks)
	},
}
