/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"strings"

	template "github.com/brayden-ooi/goo/internal/template"
	"github.com/brayden-ooi/goo/internal/utils"
	"github.com/spf13/cobra"
)

// TemplateCmd represents the now command
var TemplateCmd = &cobra.Command{
	Use:   "now",
	Short: "Scaffolds a Go project",
	Long: `Scaffolds a Go project given the size (sm | lg) according to the community recommended standards.
Refer to the link below for more information 
https://github.com/golang-standards/project-layout`,
	Run: func(cmd *cobra.Command, args []string) {
		// validation, either name or init must be populated
		if *ProjectName == "" && *ProjectInit == "" {
			log.Fatal("invalid arguments. Please supply either a name or an init")
		}

		var size string

		if *ProjectInit != "" {
			size = "lg"

			// add a default name if not supplied
			if *ProjectName == "" {
				temp := strings.Split(*ProjectInit, "/")
				*ProjectName = temp[len(temp)-1]
			}
		} else {
			size = "sm"
		}

		handler, err := template.NewHandler(
			*ProjectName,
			size,
			*TemplatePath,
		)
		if err != nil {
			log.Fatal(err)
		}

		// additional validation and operations based on size
		switch handler.Size {
		case "sm":
			err = handler.SM()

		case "lg":
			if *ProjectInit == "" {
				log.Fatal("init validation failed: Please provide a valid path for `go mod init`")
			}
			err = handler.LG(*ProjectInit)
		}

		if err != nil {
			log.Fatal(err)
		}
	},
}

var ProjectName *string  // required for sm projects, lg projects can supply one to override the default name
var ProjectInit *string  // required for lg projects
var TemplatePath *string // custom paths to templates

func init() {
	// grab installed asset path
	defaultPath, err := utils.GetDefaultTemplatePath()
	if err != nil {
		log.Fatal(err)
	}

	ProjectName = TemplateCmd.Flags().StringP("name", "n", "", "Name for the project")
	ProjectInit = TemplateCmd.Flags().StringP("init", "i", "", "Repo path for the project. Used in `go mod init` and only required for lg projects")
	TemplatePath = TemplateCmd.Flags().StringP("tmp", "t", defaultPath, "Template path for the project. Should consist of a `template-lg` and `template-sm` subdirectories. Default: ./goo")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
