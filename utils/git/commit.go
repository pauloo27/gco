package git

import (
	"os"
	"os/exec"
)

func CommitToStdout(commit string) error {
	cmd := exec.Command("git", "commit", "-m", commit)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
