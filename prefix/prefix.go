package prefix

import "strings"

type Prefix struct {
	Key, Value, Description string
}

type PrefixPack struct {
	Name, Description string
	Separator         string
	Prefixes          []*Prefix
}

func (p *PrefixPack) GetPrefix(key string) string {
	for _, prefix := range p.Prefixes {
		if strings.EqualFold(prefix.Key, key) {
			return prefix.Value + p.Separator
		}
	}
	return ""
}

var Packs []PrefixPack

func GetPrefixPack(name string) *PrefixPack {
	for _, prefixPack := range Packs {
		if strings.EqualFold(prefixPack.Name, name) {
			return &prefixPack
		}
	}
	return nil
}
