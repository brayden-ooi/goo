package create

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/brayden-ooi/goo/internal/utils"
)

func (handler Handler) SM() error {
	getTemplatePath := utils.GetProjectPath(handler.TemplatePath)
	getPath := utils.GetProjectPath(handler.Name)

	if err := utils.Copy(getTemplatePath("template-sm"), handler.Name); err != nil {
		return err
	}

	// change the permission
	err := os.Chmod(handler.Name, 0755)
	if err != nil {
		return err
	}

	// create .gitignore
	err = utils.Copy(
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
