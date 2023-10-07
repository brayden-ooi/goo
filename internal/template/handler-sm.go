package create

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/brayden-ooi/goo/internal/utils"
)

type Handler struct {
	Name         string
	Size         string
	TemplatePath string
}

func NewHandler(name, size, templatePath string) (Handler, error) {
	// Check if the name is empty or contains invalid characters
	if name == "" || strings.ContainsAny(name, "/\\:*?\"<>|") {
		return Handler{}, errors.New("invalid project name")
	}

	// size validation
	switch size {
	case "sm":
	case "lg":
	default:
		return Handler{}, errors.New("invalid size argument. Available options: sm | lg")
	}

	// template path validation
	err := validateTemplatePath(templatePath)
	if err != nil {
		return Handler{}, err
	}

	return Handler{
		Name:         name,
		Size:         size,
		TemplatePath: templatePath,
	}, nil
}

func (handler Handler) SM() error {
	getTemplatePath := utils.GetProjectPath(handler.TemplatePath)
	getPath := utils.GetProjectPath(handler.Name)

	if err := utils.Copy(getTemplatePath("template-sm"), handler.Name); err != nil {
		return err
	}

	// create .gitignore
	err := utils.Copy(
		getTemplatePath(".gitignore"),
		getPath(".gitignore"))
	if err != nil {
		return err
	}

	// initialize Git
	err = os.Chdir(getPath(""))
	if err != nil {
		return fmt.Errorf("cd operation failed: %v", err)
	}

	if err := exec.Command("git", "init").Run(); err != nil {
		return err
	}

	return nil
}

func validateTemplatePath(templatePath string) error {
	checkTemplateExist := func(path string) error {
		isTemplateExist, err := utils.CheckPathExist(path)

		if !isTemplateExist {
			return fmt.Errorf("invalid template structure: %v", path)
		}

		if err != nil {
			return fmt.Errorf("validate %s operation failed: %v", path, err)
		}

		return nil
	}

	getPath := utils.GetProjectPath(templatePath)
	// check if there is template-lg
	err := checkTemplateExist(getPath("template-lg"))
	if err != nil {
		return err
	}

	// check if there is template-sm
	err = checkTemplateExist(getPath("template-sm"))
	if err != nil {
		return err
	}

	// check if there is .gitignore
	err = checkTemplateExist(getPath(".gitignore"))
	if err != nil {
		return err
	}

	return nil
}
