package mode

import (
	"fmt"

	"github.com/c-bata/go-prompt"
)

func initCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "text", Description: "feat: fix: ci:"},
		{Text: "emojis-name", Description: ":sparkles: :bug: :construction_worker:"},
		{Text: "emojis-unicode", Description: "✨ 🐛 👷"},
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
