package mode

import (
	"fmt"
	"os"

	"github.com/Pauloo27/gommit/config"
	"github.com/Pauloo27/gommit/prefix"
	"github.com/Pauloo27/gommit/utils"
	"github.com/c-bata/go-prompt"
)

func commitCompleter(prefixPack *prefix.PrefixPack) prompt.Completer {
	return func(d prompt.Document) []prompt.Suggest {
		s := []prompt.Suggest{}
		for _, prefix := range prefixPack.Prefixes {
			s = append(s, prompt.Suggest{
				Text:        prefix.Value,
				Description: prefix.Name,
			})
		}
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

	prefix := utils.Prompt(" λ ", commitCompleter(pack), prompt.OptionPrefixTextColor(prompt.Blue))
	prefix = pack.GetPrefix(prefix)

	// TODO: prompt title
	utils.MoveCursorUp(1)
	utils.ClearLine()
	title := utils.Prompt(" λ "+prefix, utils.EmptyCompleter, prompt.OptionPrefixTextColor(prompt.Blue))
	fmt.Println(prefix + title)
	// TODO: prompt message
}
