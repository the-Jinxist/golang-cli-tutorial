/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "golang-cli-tutorial",
	Short: "A brief description of your application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var shout = &cobra.Command{
	Use:     "shout",
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
	rootCmd.AddCommand(taskCommand)
	rootCmd.AddCommand(shout)
}
