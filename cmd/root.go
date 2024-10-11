/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
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
	Use:   "golang-cli-tutorial",
	Short: "A brief description of your application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var shout = &cobra.Command{
	Use:     "ping",
	Short:   "Tell the cli to shout your name",
	Example: "golang-cli-tutorial shout Favour",
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

	db := config.GetDB()

	repo := GetRepo(db)

	taskCmd := getTaskCommand(repo)
	addTaskCmd := addTaskCommand(repo)

	taskCmd.Flags().StringP("project", "p", "", "specify the project of this task")

	addTaskCmd.Flags().StringP("name", "n", "", "specify the name of your task (required)")
	addTaskCmd.MarkFlagRequired("name")
	addTaskCmd.Flags().StringP("project", "p", "", "specify the project of this task")

	rootCmd.AddCommand(taskCmd)
	rootCmd.AddCommand(addTaskCmd)
	rootCmd.AddCommand(shout)
}
