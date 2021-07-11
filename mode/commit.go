package mode

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/Pauloo27/gommit/config"
	"github.com/Pauloo27/gommit/prefix"
	"github.com/Pauloo27/gommit/utils"
	"github.com/c-bata/go-prompt"
)

func commitCompleter(prefixPack *prefix.PrefixPack) prompt.Completer {
	s := []prompt.Suggest{}
	for _, prefix := range prefixPack.Prefixes {
		s = append(s, prompt.Suggest{
			Text:        prefix.Name,
			Description: prefix.Description,
		})
	}
	return func(d prompt.Document) []prompt.Suggest {
		return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
	}
}

func Commit() {
	c, err := config.GetProjectConfig()
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Gommit config not found. Create one with gommit --init")
			os.Exit(-1)
		}
	}
	pack := prefix.GetPrefixPack(c.PrefixPack)
	if pack == nil {
		fmt.Println("Prefix pack not found")
		os.Exit(-1)
	}
	// TODO: prompt files to add

	promptPrefix := " -> "

	fmt.Println("Enter a empty line to cancel the commit")
	fmt.Printf("%s%s\n", strings.Repeat(" ", len(promptPrefix)), strings.Repeat("-", 49))

	rawPrefix := utils.Prompt(promptPrefix, commitCompleter(pack), prompt.OptionPrefixTextColor(prompt.Blue))
	prefix := pack.GetPrefix(rawPrefix)
	if prefix == "" {
		prefix = rawPrefix
	}

	utils.MoveCursorUp(1)
	utils.ClearLine()
	title := utils.Prompt(promptPrefix+prefix, utils.EmptyCompleter, prompt.OptionPrefixTextColor(prompt.Blue))
	if title == "" {
		fmt.Println("Commit cancelled")
		os.Exit(-1)
	}
	message := ""
	fmt.Println(" == Write the commit body, line by line")
	fmt.Println(" == Enter a line with spaces to add a empty line")
	fmt.Println(" == Enter a empty line when you are done")
	fmt.Printf("%s%s\n", strings.Repeat(" ", len(promptPrefix)), strings.Repeat("-", 82))
	for {
		line := prompt.Input(promptPrefix, utils.EmptyCompleter)
		if line == "" {
			break
		}
		if strings.TrimSpace(line) == "" {
		}
		message += line
	}
	commit := prefix + title + "\n\n" + message
	cmd := exec.Command("git", "commit", "-m", commit)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		os.Exit(-1)
	}
	fmt.Println("nice")
}
