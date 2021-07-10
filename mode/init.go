package mode

import (
	"fmt"

	"github.com/Pauloo27/gommit/prefix"
	"github.com/c-bata/go-prompt"
)

func initCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{}
	for _, prefixPack := range prefix.Packs {
		s = append(s, prompt.Suggest{
			Text: prefixPack.Name,
			Description: fmt.Sprintf("%s %s %s",
				prefixPack.GetPrefix("feat"), prefixPack.GetPrefix("fix"), prefixPack.GetPrefix("ci"),
			),
		})
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func Init() {
	name := prompt.Input(
		"Prefix pack: ", initCompleter,
		prompt.OptionDescriptionTextColor(prompt.White),
		prompt.OptionDescriptionBGColor(prompt.Black),

		prompt.OptionSuggestionTextColor(prompt.White),
		prompt.OptionSuggestionBGColor(prompt.Black),

		prompt.OptionSelectedSuggestionTextColor(prompt.LightGray),
		prompt.OptionSelectedSuggestionBGColor(prompt.Black),

		prompt.OptionSelectedDescriptionTextColor(prompt.Black),
		prompt.OptionSelectedDescriptionBGColor(prompt.LightGray),

		prompt.OptionShowCompletionAtStart(),
	)
	fmt.Println(name)
}
