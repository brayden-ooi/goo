package create

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
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

func handlePermissions(path string) error {
	err := filepath.WalkDir(path, func(path string, di fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		err = os.Chmod(path, 0755)

		return err
	})

	if err != nil {
		return fmt.Errorf("could not chmod recursively: %v", path)
	}

	return nil
}

func HandleGitIgnore(src, dst string) error {
	err := utils.Copy(src, dst)
	if err != nil {
		return fmt.Errorf("copy .gitignore operation failed: %v", err)
	}

	err = os.Chmod(dst, 0755)
	if err != nil {
		return fmt.Errorf("update .gitignore permission operation failed: %v", err)
	}

	return nil
}
