/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goo",
	Short: "A set of utilities to get started with Go",
	Long:  `Goo is a set of utilities that scaffolds or helps with managing a Go project.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addCommands() {
	rootCmd.AddCommand(TemplateCmd)
}

func init() {
	addCommands()
}
