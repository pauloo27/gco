package utils

import "github.com/c-bata/go-prompt"

func Prompt(prefix string, completer prompt.Completer, options ...prompt.Option) string {
	opts := []prompt.Option{
		prompt.OptionDescriptionTextColor(prompt.White),
		prompt.OptionDescriptionBGColor(prompt.Black),

		prompt.OptionSuggestionTextColor(prompt.White),
		prompt.OptionSuggestionBGColor(prompt.Black),

		prompt.OptionSelectedSuggestionTextColor(prompt.LightGray),
		prompt.OptionSelectedSuggestionBGColor(prompt.Black),

		prompt.OptionSelectedDescriptionTextColor(prompt.Black),
		prompt.OptionSelectedDescriptionBGColor(prompt.LightGray),

		prompt.OptionShowCompletionAtStart(),
	}
	opts = append(opts, options...)
	return prompt.Input(
		prefix, completer,
		opts...,
	)
}

func EmptyCompleter(d prompt.Document) []prompt.Suggest {
	return nil
}
