package mode

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/Pauloo27/gco/config/holder"
	"github.com/Pauloo27/gco/prefix"
	"github.com/Pauloo27/gco/utils"
	"github.com/Pauloo27/gco/utils/git"
	"github.com/c-bata/go-prompt"
)

func commitCompleter(prefixPack *prefix.PrefixPack) prompt.Completer {
	s := []prompt.Suggest{}
	for _, prefix := range prefixPack.Prefixes {
		s = append(s, prompt.Suggest{
			Text:        prefix.Value,
			Description: prefix.Description,
		})
	}
	return func(d prompt.Document) []prompt.Suggest {
		return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
	}
}

func listChangedFiles(out prompt.ConsoleWriter) ([]*git.ChangedFile, []string) {
	files, err := git.GetChangedFiles()

	if err != nil {
		fmt.Println("Cannot get repository status")
		os.Exit(-1)
	}

	if len(files) == 0 {
		utils.PrettyPrint(out, "Nothing changed since last commit\n")
		os.Exit(-1)
	}

	fmt.Println()
	utils.PrettyPrint(out, "Repository status:\n")
	utils.PrettyPrint(out, "(", prompt.Green, "green ", prompt.DefaultColor,
		"files are going to committed)\n",
	)
	filesName := []string{}
	for id, file := range files {
		color := prompt.Red
		if file.Tracked {
			color = prompt.Green
		}
		utils.PrettyPrint(out,
			" -> ", color, id+1, " ", file.Name, file.Status, prompt.DefaultColor, "\n",
		)
		filesName = append(filesName, file.Name)
	}
	fmt.Println()
	return files, filesName
}

func Commit(skipHooks bool) {
	conf, isProjectConf, err := holder.GetProjectConfigOrGlobal()
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("GCO config not found. Create one with gco --init")
			fmt.Println("You can also create a global config with gco --init-global")
			os.Exit(-1)
		}
		fmt.Println(err)
		os.Exit(-1)
	}

	if !isProjectConf {
		fmt.Println("Using global config, to create one for you project, run gco --init")
	}

	pack := prefix.GetPrefixPack(conf.PrefixPack)
	if pack == nil {
		fmt.Println("Prefix pack not found")
		os.Exit(-1)
	}

	out := prompt.NewStderrWriter()

	branch, err := git.GetCurrentBranchName()
	if err != nil {
		fmt.Println("Cannot get current branch name")
		os.Exit(-1)
	}
	utils.PrettyPrint(out,
		"You are committing to ", prompt.Blue, branch, prompt.DefaultColor, "\n",
	)

	files, filesName := listChangedFiles(out)

	if conf.AskFilesToCommit {
		err = git.PromptFilesToAdd(out, filesName, files)
		if err != nil {
			fmt.Println("Cannot add files: ", err)
			os.Exit(-1)
		}
	}

	if !skipHooks {
		git.CallPreCommitHooks(conf)
	}

	promptPrefix := " λ "
	promptPrefixLen := utf8.RuneCountInString(promptPrefix)

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
		fmt.Println("Something went wrong while writing to output console")
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

	if conf.AskCISkip {
		input := utils.PromptYesOrNot(
			" Skip CI? [y/N] ",
		)
		if input == "yes" {
			message += "[ci skip]"
		}
	}

	commit := prefix + title + "\n\n" + message
	err = git.CommitToStdout(commit)
	if err != nil {
		err := os.WriteFile(".gcobkp", []byte(commit), 0600)
		if err != nil {
			os.Exit(-1)
		}
		utils.PrettyPrint(out,
			"Looks like the commit ", prompt.Red, "failed", prompt.DefaultColor,
			". No worries! It was saved to ", prompt.Blue, ".gcobkp",
			prompt.DefaultColor, "\n",
		)
		utils.PrettyPrint(out,
			"You can ", prompt.Blue, "retry ", prompt.DefaultColor, "using ",
			prompt.Blue, "gco --retry", prompt.DefaultColor, "\n",
		)
		os.Exit(-1)
	}

	if !skipHooks {
		git.CallPostCommitHooks(conf)
	}
}
