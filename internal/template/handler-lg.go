package create

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/brayden-ooi/goo/internal/utils"
)

func (handler Handler) LG(init string) error {
	getTemplatePath := utils.GetProjectPath(handler.TemplatePath)
	getPath := utils.GetProjectPath(handler.Name)

	err := utils.Copy(getTemplatePath("template-lg"), handler.Name)
	if err != nil {
		return err
	}

	// change the permission
	err = os.Chmod(handler.Name, 0755)
	if err != nil {
		return err
	}

	// additional steps to perform
	// rename main.txt to main.go
	err = os.Rename(getPath("main.txt"), getPath("main.go"))
	if err != nil {
		return err
	}

	// replace $PROJECT_REPO
	err = utils.ReplaceText(getPath("main.go"), "$PROJECT_REPO", init)
	if err != nil {
		return err
	}

	// replace $PROJECT_NAME
	err = utils.ReplaceText(getPath("Makefile"), "$PROJECT_NAME", handler.Name)
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

	// exec go mod init
	err = os.Chdir(getPath(""))
	if err != nil {
		return fmt.Errorf("cd operation failed: %v", err)
	}

	err = exec.Command("go", "mod", "init", init).Run()
	if err != nil {
		return err
	}

	// git init
	err = exec.Command("git", "init").Run()
	if err != nil {
		return err
	}

	return nil
}
