package utils

import (
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
	return RunGitCommand("rev-parse", "--show-toplevel")
}

func GetCurrentBranchName() (string, error) {
	return RunGitCommand("rev-parse", "--abbrev-ref", "HEAD")
}
