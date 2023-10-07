package utils

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// https://opensource.com/article/18/6/copying-files-go
func Copy(src, dst string) error {
	// if dst already exist, throw
	if _, err := os.Stat(dst); err == nil {
		return fmt.Errorf("path `%s` already exist", dst)
	} else if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
	}

	// execute cp recursively from src to dst
	if err := exec.Command("cp", "-r", src, dst).Run(); err != nil {
		return err
	}

	return nil
}

// git logic handler
func HandleGit() error {
	// initialize Git
	if err := exec.Command("git", "init").Run(); err != nil {
		return err
	}

	// // create .gitignore
	// f, err := os.Create(".gitignore")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()

	// // write into
	// f.Write()

	return nil
}

// https://stackoverflow.com/questions/26152901/replace-a-line-in-text-file-golang
func ReplaceText(fileName, prevTxt, nextTxt string) error {
	input, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, prevTxt) {
			lines[i] = nextTxt
		}
	}

	output := strings.Join(lines, "\n")

	err = os.WriteFile(fileName, []byte(output), 0644)
	if err != nil {
		return err
	}

	return nil
}

func GetProjectPath(projectName string) func(string) string {
	return func(fileName string) string {
		return fmt.Sprintf("%s/%s", projectName, fileName)
	}
}
