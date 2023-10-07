package create

import (
	"os"
	"os/exec"

	"github.com/brayden-ooi/goo/internal/utils"
)

func HandlerLG(dst string, init string) error {
	err := utils.Copy("assets/template-lg", dst)
	if err != nil {
		return err
	}

	getPath := utils.GetProjectPath(dst)

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
	err = utils.ReplaceText(getPath("main.go"), "$PROJECT_NAME", dst)
	if err != nil {
		return err
	}

	// exec go mod init
	err = exec.Command("go", "mod", "init", init).Run()
	if err != nil {
		return err
	}

	// if err := utils.HandleGit(); err != nil {
	// 	return err
	// }

	return nil
}
