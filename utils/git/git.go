package git

import (
	"errors"
	"os/exec"
	"strings"
)

func RunGitCommand(params ...string) (string, error) {
	cmd := exec.Command("git", params...)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(string(out), "\n"), nil
}

func GetRepositoryRoot() (string, error) {
	root, err := RunGitCommand("rev-parse", "--show-toplevel")
	if err != nil {
		err = errors.New("You are not inside a git repository")
	}
	return root, err
}
