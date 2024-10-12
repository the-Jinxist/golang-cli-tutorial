/*
Copyright © 2024 Favour Olukayode <nerosilva522@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/the-Jinxist/golang-cli-tutorial/config"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "todo",
	Aliases: []string{"todo"},
	Short:   "A brief description of your application",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var shout = &cobra.Command{
	Use:     "ping",
	Short:   "Tell the cli to shout your name",
	Example: "todo shout Favour",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s!!!!!!!!!!", strings.ToUpper(args[0]))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(shout)

	db := config.GetDB()
	repo := GetRepo(db)

	taskCmd := getTaskCommand(repo)
	addTaskCmd := addTaskCommand(repo)
	finishTaskCmd := finishTaskCommand(repo)
	deleteTaskCmd := deleteTaskCommand(repo)
	startTaskCmd := startTaskCommand(repo)
	clearTaskCmd := clearTaskCoomands(repo)

	taskCmd.Flags().StringP("project", "p", "", "specify the project of this task")
	taskCmd.Flags().StringP("status", "s", "", "filter project by status")
	rootCmd.AddCommand(taskCmd)

	addTaskCmd.Flags().StringP("project", "p", "", "specify the project of this task")
	rootCmd.AddCommand(addTaskCmd)

	finishTaskCmd.Flags().IntP("id", "i", -1, "specify the id of the task (required)")
	finishTaskCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(finishTaskCmd)

	deleteTaskCmd.Flags().IntP("id", "i", -1, "specify the id of the task (required)")
	deleteTaskCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(deleteTaskCmd)

	startTaskCmd.Flags().IntP("id", "i", -1, "specify the id of the task (required)")
	startTaskCmd.MarkFlagRequired("id")
	rootCmd.AddCommand(startTaskCmd)

	rootCmd.AddCommand(clearTaskCmd)
}
