package mode

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/Pauloo27/gommit/config"
	"github.com/Pauloo27/gommit/prefix"
	"github.com/Pauloo27/gommit/utils"
	"github.com/Pauloo27/gommit/utils/git"
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

	out := prompt.NewStderrWriter()
	promptPrefix := " λ "
	promptPrefixLen := utf8.RuneCountInString(promptPrefix)

	branch, err := git.GetCurrentBranchName()
	utils.PrettyPrint(out,
		"You are commiting to ", prompt.Blue, branch, prompt.DefaultColor, "\n",
	)

	fmt.Println("Enter a empty line to cancel the commit")
	fmt.Printf("%s%s\n", strings.Repeat(" ", promptPrefixLen), strings.Repeat("-", 49))

	rawPrefix := utils.Prompt(promptPrefix, commitCompleter(pack), prompt.OptionPrefixTextColor(prompt.Blue))
	if rawPrefix == "" {
		fmt.Println("Commit cancelled")
		os.Exit(-1)
	}
	prefix := pack.GetPrefix(rawPrefix)
	if prefix == "" {
		prefix = rawPrefix
	}

	out.CursorUp(1)
	out.EraseLine()
	err = out.Flush()
	if err != nil {
		fmt.Println("Something went wrong while writting to output console")
		os.Exit(-1)
	}

	title := utils.Prompt(promptPrefix+prefix, utils.EmptyCompleter, prompt.OptionPrefixTextColor(prompt.Blue))
	if title == "" {
		fmt.Println("Commit cancelled")
		os.Exit(-1)
	}
	message := ""

	utils.PrettyPrint(out, " == Write the commit body, line by line\n")
	utils.PrettyPrint(out,
		" == Enter a line ", prompt.Blue, "with spaces ", prompt.DefaultColor,
		"to add a ", prompt.Blue, "empty line", prompt.DefaultColor, "\n",
	)
	utils.PrettyPrint(out,
		" == Enter an ", prompt.Blue, "empty line ", prompt.DefaultColor,
		"when you are done\n",
	)
	utils.PrettyPrint(out,
		" == Enter a ", prompt.Blue, ". ", prompt.DefaultColor,
		"to ", prompt.Blue, "cancel commit", prompt.DefaultColor, "\n",
	)

	fmt.Printf("%s%s\n", strings.Repeat(" ", promptPrefixLen), strings.Repeat("-", 82))
	for {
		line := prompt.Input(promptPrefix, utils.EmptyCompleter)
		if line == "." {
			fmt.Println("Commit cancelled")
			os.Exit(-1)
		}
		if line == "" {
			break
		}
		if strings.TrimSpace(line) != "" {
			message += line
		}
		message += "\n"
	}
	err = git.CommitToStdout(prefix, title, message)
	if err != nil {
		os.Exit(-1)
	}
}
