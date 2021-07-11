package git

import (
	"os"
	"os/exec"
)

func CommitToStdout(prefix, title, message string) error {
	commit := prefix + title + "\n\n" + message
	cmd := exec.Command("git", "commit", "-m", commit)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
