package git

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/Pauloo27/gco/config"
	"github.com/Pauloo27/gco/utils"
	"github.com/c-bata/go-prompt"
)

func callHooks(title string, hooks []config.HookAction) {
	if len(hooks) == 0 {
		return
	}
	out := prompt.NewStdoutWriter()
	utils.PrettyPrint(
		out,
		"Running ", prompt.Blue, len(hooks), prompt.DefaultColor,
		" ", title, " hooks\n",
	)
	for _, hook := range hooks {
		utils.PrettyPrint(
			out,
			"-> ", prompt.Yellow, hook.Command, prompt.DefaultColor, "\n",
		)
		if hook.Ask {
			skip := utils.PromptYesOrNot(" Skip? [y/N] ")
			if skip == "yes" {
				utils.PrettyPrint(
					out,
					prompt.Green, " Skipped", prompt.DefaultColor, "\n",
				)
				continue
			}
		}
		rawCommand := strings.Split(hook.Command, " ")
		cmd := exec.Command(rawCommand[0], rawCommand[1:]...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil && hook.ExitOnFail {
			fmt.Println("A required hook failed")
			os.Exit(-1)
		}
	}
}

func CallPreCommitHooks(conf *config.Config) {
	callHooks("pre-commit", conf.PreCommit)
}

func CallPostCommitHooks(conf *config.Config) {
	callHooks("post-commit", conf.PostCommit)
}

func CommitToStdout(commit string) error {
	cmd := exec.Command("git", "commit", "-m", commit)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
