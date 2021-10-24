package mode

import (
	"fmt"
	"os"

	"github.com/Pauloo27/gco/internal/config"
	"github.com/Pauloo27/gco/internal/config/holder"
	"github.com/Pauloo27/gco/internal/prefix"
	"github.com/Pauloo27/gco/internal/utils"
	"github.com/Pauloo27/gco/internal/version"
	"github.com/c-bata/go-prompt"
)

func initCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{}
	for _, prefixPack := range prefix.Packs {
		s = append(s, prompt.Suggest{
			Text: prefixPack.Name,
			Description: fmt.Sprintf("%s %s %s",
				prefixPack.GetPrefix("feat").Value, prefixPack.GetPrefix("fix").Value, prefixPack.GetPrefix("ci").Value,
			),
		})
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func initConf(isProject bool) {
	name := utils.Prompt("Prefix pack: ", initCompleter)

	pack := prefix.GetPrefixPack(name)
	if pack == nil {
		fmt.Println("Prefix pack not found =(")
		os.Exit(-1)
	}
	c := config.Config{
		PrefixPack: pack.Name,
		Version:    version.VERSION,
	}
	var err error
	if isProject {
		err = holder.StoreProjectConfig(&c)
	} else {
		err = holder.StoreGlobalConfig(&c)
	}
	if err != nil {
		fmt.Println("Something went wrong while storing the config", err)
		os.Exit(-1)
	}
}

func Init() {
	initConf(true)
}

func InitGlobal() {
	initConf(false)
}
