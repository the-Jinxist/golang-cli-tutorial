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
			status, _ := cmd.Flags().GetString("status")

			if project != "" && len(project) <= 3 {
				return fmt.Errorf("project name must be more than 3 chars")
			}

			if status != "" && len(status) <= 3 {
				return fmt.Errorf("status name must be more than 3 chars")
			}
			tasks, err := repo.getTasks(project, status)
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
		Args:  cobra.MaximumNArgs(20),
		RunE: func(cmd *cobra.Command, args []string) error {

			project, _ := cmd.Flags().GetString("project")
			if project != "" && len(project) <= 3 {
				return fmt.Errorf("project name must be more than 3")
			}

			//TODO: ADD SUPPORT FOR MUTLIPLE LINE TASK NAMES
			name := args[0]
			task := Task{
				Name:    name,
				Project: project,
				Status:  "pending",
			}

			err := repo.createTask(task)
			if err != nil {
				return err
			}

			fmt.Print("task added successfully\n")

			return nil
		},
	}
	return taskCommand
}

func finishTaskCommand(repo repository) *cobra.Command {
	taskCommand := &cobra.Command{
		Use:   "finish",
		Short: "This is command is used to mark a task as finished",
		RunE: func(cmd *cobra.Command, args []string) error {

			id, _ := cmd.Flags().GetInt("id")
			if id < 0 {
				return fmt.Errorf("invalid")
			}
			taskName, err := repo.finishTask(id)
			if err != nil {
				return err
			}

			fmt.Printf("%s is marked as finished!\n", taskName)

			return nil
		},
	}
	return taskCommand
}

func deleteTaskCommand(repo repository) *cobra.Command {
	taskCommand := &cobra.Command{
		Use:   "delete",
		Short: "This is command is used to mark a task as finished",
		RunE: func(cmd *cobra.Command, args []string) error {

			id, _ := cmd.Flags().GetInt("id")
			if id < 0 {
				return fmt.Errorf("invalid")
			}
			taskName, err := repo.deleteTask(id)
			if err != nil {
				return err
			}

			fmt.Printf("%s is deleted!\n", taskName)

			return nil
		},
	}
	return taskCommand
}

func startTaskCommand(repo repository) *cobra.Command {
	taskCommand := &cobra.Command{
		Use:   "start",
		Short: "This is command is used to mark a task as in progress",
		RunE: func(cmd *cobra.Command, args []string) error {

			id, _ := cmd.Flags().GetInt("id")
			if id < 0 {
				return fmt.Errorf("invalid")
			}
			taskName, err := repo.startTask(id)
			if err != nil {
				return err
			}

			fmt.Printf("%s is started!\n", taskName)

			return nil
		},
	}
	return taskCommand
}

func clearTaskCoomands(repo repository) *cobra.Command {
	taskCommand := &cobra.Command{
		Use:   "clear_all",
		Short: "This is command is used to mark a task as finished",
		RunE: func(cmd *cobra.Command, args []string) error {

			err := repo.clearAllTasks()
			if err != nil {
				return err
			}

			fmt.Printf("All tasks have been cleared successfully!\n")

			return nil
		},
	}
	return taskCommand
}
