package utils

import (
	"errors"
	"fmt"
	"go/build"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// https://opensource.com/article/18/6/copying-files-go
func Copy(src, dst string) error {
	// if dst already exist, throw
	isFileExist, err := CheckPathExist(dst)
	if isFileExist {
		return fmt.Errorf("path `%s` already exist", dst)
	}
	if err != nil {
		return err
	}

	// execute cp recursively from src to dst
	cpCmd := exec.Command("cp", "-r", src, dst)
	if err := cpCmd.Run(); err != nil {
		return fmt.Errorf("copy operation failed: %v", err)
	}

	return nil
}

func CheckPathExist(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	} else {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		} else {
			return false, err
		}
	}
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
			lines[i] = strings.ReplaceAll(line, prevTxt, nextTxt)
		}
	}

	output := strings.Join(lines, "\n")

	err = os.WriteFile(fileName, []byte(output), 0755)
	if err != nil {
		return fmt.Errorf("replace text operation failed: %v", err)
	}

	return nil
}

func GetProjectPath(projectName string) func(string) string {
	return func(fileName string) string {
		return fmt.Sprintf("%s/%s", projectName, fileName)
	}
}

func GetDefaultTemplatePath() (string, error) {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	grandparent := filepath.Join(gopath, "pkg", "mod", "github.com", "brayden-ooi")
	entries, err := os.ReadDir(grandparent)
	if err != nil {
		return "", err
	}
	var parent string

	for _, dirname := range entries {
		if strings.ContainsAny(dirname.Name(), "goo") {
			parent = dirname.Name()
		}
	}

	if parent == "" {
		return "", errors.New("could not retrieve default template path")
	}

	return filepath.Join(grandparent, parent, "assets"), nil
}
