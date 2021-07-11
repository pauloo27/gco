package prefix

import "strings"

type Prefix struct {
	Name, Value string
}

type PrefixPack struct {
	Name, Description string
	Separator         string
	Prefixes          []*Prefix
}

func (p *PrefixPack) GetPrefix(name string) string {
	for _, prefix := range p.Prefixes {
		if strings.EqualFold(prefix.Value, name) {
			return prefix.Value + p.Separator
		}
	}
	return ""
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
