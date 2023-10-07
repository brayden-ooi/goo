package create

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/brayden-ooi/goo/internal/utils"
)

func (handler Handler) LG(init string) error {
	err := utils.Copy(utils.GetProjectPath(handler.TemplatePath)("template-lg"), handler.Name)
	if err != nil {
		return err
	}

	getPath := utils.GetProjectPath(handler.Name)

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

	// exec go mod init
	err = os.Chdir(getPath(""))
	if err != nil {
		return fmt.Errorf("cd operation failed: %v", err)
	}

	initCmd := exec.Command("go", "mod", "init", init)
	err = initCmd.Run()
	if err != nil {
		return err
	}

	// if err := utils.HandleGit(); err != nil {
	// 	return err
	// }

	return nil
}
