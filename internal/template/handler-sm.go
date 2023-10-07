package create

import (
	"github.com/brayden-ooi/goo/internal/utils"
)

func HandlerSM(dst string) error {
	if err := utils.Copy("assets/template-sm", dst); err != nil {
		return err
	}

	if err := utils.HandleGit(); err != nil {
		return err
	}

	return nil
}
