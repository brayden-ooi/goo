/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"log"
	"strings"

	template "github.com/brayden-ooi/goo/internal/template"
	"github.com/spf13/cobra"
)

func validateName(input *string) (string, error) {
	// Check if the name is empty or contains invalid characters
	if *input == "" || strings.ContainsAny(*input, "/\\:*?\"<>|") {
		return "", errors.New("invalid project name")
	}

	return *input, nil
}

func validateSize(input *string) (string, error) {
	switch *input {
	case "sm":
		return "sm", nil
	case "lg":
		return "lg", nil
	default:
		return "", errors.New("invalid size argument. Available options: sm | lg")
	}
}

// TemplateCmd represents the now command
var TemplateCmd = &cobra.Command{
	Use:   "now",
	Short: "Scaffolds a Go project",
	Long: `Scaffolds a Go project given the size (sm | lg) according to the community recommended standards.
Refer to the link below for more information 
https://github.com/golang-standards/project-layout`,
	Run: func(cmd *cobra.Command, args []string) {
		name, err := validateName(ProjectName)
		if err != nil {
			log.Fatal("something went wrong: ", err)
		}

		size, err := validateSize(ProjectSize)
		if err != nil {
			log.Fatal("something went wrong: ", err)
		}

		// additional validation and operations based on size
		switch size {
		case "sm":
			err = template.HandlerSM(name)

		case "lg":
			// go mod init will handle the validation and exiting for invalid ProjectInit
			err = template.HandlerLG(name, *ProjectInit)
		}

		if err != nil {
			log.Fatal("something went wrong: ", err)
		}
	},
}

var ProjectName *string // required
var ProjectSize *string // defaults to sm
var ProjectInit *string // only required for lg projects

func init() {
	ProjectName = TemplateCmd.Flags().StringP("name", "n", "", "Name for the project")
	ProjectSize = TemplateCmd.Flags().StringP("size", "s", "sm", "Preset templates to generate. Available options: sm | lg")
	ProjectInit = TemplateCmd.Flags().StringP("init", "i", "", "Repo path for the project. Used in `go mod init` and only required for lg projects")

	// required
	if err := TemplateCmd.MarkFlagRequired("name"); err != nil {
		fmt.Println(err)
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
