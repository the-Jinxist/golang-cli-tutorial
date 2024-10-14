package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func getTaskCommand(repo repository) *cobra.Command {
	taskCommand := &cobra.Command{
		Use:     "tasks",
		Short:   "This is a list of all your tasks. You can filter by project with the \"project\" flag",
		Example: "golang-cli-tutorial tasks -s <status> -p <project>",
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
		Use:     "add",
		Short:   "This is command is used to add tasks",
		Example: "golang-cli-tutorial add <task name> -p <project>",
		Args:    cobra.MaximumNArgs(20),
		RunE: func(cmd *cobra.Command, args []string) error {

			project, _ := cmd.Flags().GetString("project")
			if project != "" && len(project) <= 3 {
				return fmt.Errorf("project name must be more than 3")
			}

			//TODO: ADD SUPPORT FOR MUTLIPLE LINE TASK NAMES
			name := strings.Join(args, " ")
			task := Task{
				Name:    name,
				Project: project,
				Status:  "pending",
			}

			err := repo.createTask(task)
			if err != nil {
				return err
			}

			res := bannerRes("task added successfully", "", "")
			fmt.Println(res)

			return nil
		},
	}
	return taskCommand
}

func finishTaskCommand(repo repository) *cobra.Command {
	taskCommand := &cobra.Command{
		Use:     "finish",
		Short:   "This is command is used to mark a task as finished",
		Example: "golang-cli-tutorial finish <id>",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			taskName, err := repo.finishTask(id)
			if err != nil {
				return err
			}

			res := bannerRes(fmt.Sprintf("[%s] is marked as finished!", taskName), "", "")
			fmt.Println(res)

			return nil
		},
	}
	return taskCommand
}

func deleteTaskCommand(repo repository) *cobra.Command {
	taskCommand := &cobra.Command{
		Use:     "delete",
		Short:   "This is command is used to delete a task",
		Example: "golang-cli-tutorial delete <id>",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			taskName, err := repo.deleteTask(id)
			if err != nil {
				return err
			}

			res := bannerRes(fmt.Sprintf("[%s] is deleted!", taskName), "", "#FF5733")
			fmt.Println(res)

			return nil
		},
	}
	return taskCommand
}

func startTaskCommand(repo repository) *cobra.Command {
	taskCommand := &cobra.Command{
		Use:     "start",
		Short:   "This is command is used to mark a task as in progress",
		Example: "golang-cli-tutorial start <id>",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			id, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			taskName, err := repo.startTask(id)
			if err != nil {
				return err
			}

			res := bannerRes(fmt.Sprintf("[%s] is started!", taskName), "", "#40B5AD")
			fmt.Println(res)

			return nil
		},
	}
	return taskCommand
}

func clearTaskCoomands(repo repository) *cobra.Command {
	taskCommand := &cobra.Command{
		Use:     "clear_all",
		Short:   "This is command is used to mark a task as finished",
		Example: "golang-cli-tutorial clear_all",
		RunE: func(cmd *cobra.Command, args []string) error {

			err := repo.clearAllTasks()
			if err != nil {
				return err
			}

			res := bannerRes("All tasks have been cleared successfully", "", "")
			fmt.Println(res)

			return nil
		},
	}
	return taskCommand
}
