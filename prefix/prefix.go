package prefix

import "strings"

type Prefix struct {
	Name, Value string
}

type PrefixPack struct {
	Name, Description string
	Separator         string
	Prefixes          map[string]*Prefix
}

func (p *PrefixPack) GetPrefix(name string) string {
	return p.Prefixes[name].Value + p.Separator
}

var Packs = []PrefixPack{
	TextPrefix,
}

func GetPrefixPack(name string) *PrefixPack {
	for _, prefixPack := range Packs {
		if strings.EqualFold(prefixPack.Name, name) {
			return &prefixPack
		}
	}
	return nil
}
