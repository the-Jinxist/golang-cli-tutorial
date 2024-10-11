package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func getTaskCommand(repo repository) *cobra.Command {
	taskCommand := &cobra.Command{
		Use:   "tasks",
		Short: "This is a list of all your tasks. You can filter by project with the \"project\" flag",
		RunE: func(cmd *cobra.Command, args []string) error {

			project, _ := cmd.Flags().GetString("project")
			if project != "" && len(project) <= 3 {
				return fmt.Errorf("project name must be more than 3")
			}
			tasks, err := repo.getTasks(project)
			if err != nil {
				return err
			}

			table := setupTable(tasks)
			fmt.Println(table)

			return nil
		},
	}
	return taskCommand
}

func addTaskCommand(repo repository) *cobra.Command {
	taskCommand := &cobra.Command{
		Use:   "add",
		Short: "This is command is used to add tasks",
		RunE: func(cmd *cobra.Command, args []string) error {

			project, _ := cmd.Flags().GetString("project")
			if project != "" && len(project) <= 3 {
				return fmt.Errorf("project name must be more than 3")
			}

			name, _ := cmd.Flags().GetString("name")
			task := Task{
				Name:    name,
				Project: project,
				Status:  "pending",
			}

			err := repo.createTask(task)
			if err != nil {
				return err
			}

			fmt.Print("task added successfully")

			return nil
		},
	}
	return taskCommand
}
