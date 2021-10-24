package mode

import (
	"fmt"
	"os"

	"github.com/Pauloo27/gco/internal/utils/git"
)

func Retry() {
	file, err := os.ReadFile(".gcobkp")
	if err != nil {
		fmt.Println("Cannot read file .gcobkp")
		os.Exit(-1)
	}
	commit := string(file)
	err = git.CommitToStdout(commit)
	if err != nil {
		fmt.Println("Retry failed")
		os.Exit(-1)
	}
	err = os.Remove(".gcobkp")
	if err != nil {
		fmt.Println("Cannot delete file .gcobkp")
		os.Exit(-1)
	}
}
