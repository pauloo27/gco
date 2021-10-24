package prefix

import (
	"fmt"
	"strings"
)

type Prefix struct {
	Key, Value, Description string
	pack                    *PrefixPack
}

type PrefixPack struct {
	Name, Description string
	Separator         string
	Prefixes          []*Prefix
}

func (p *PrefixPack) FormatPrefix(prefix, module string) string {
	if module == "" {
		return fmt.Sprintf("%s%s", prefix, p.Separator)
	}
	return fmt.Sprintf("%s(%s)%s", prefix, module, p.Separator)
}

func (p *PrefixPack) FormatPrefixForModulePrompt(prefix string) string {
	return fmt.Sprintf("%s(", prefix)
}

func (p *PrefixPack) GetPrefix(key string) *Prefix {
	for _, prefix := range p.Prefixes {
		if strings.EqualFold(prefix.Key, key) {
			prefix.pack = p
			return prefix
		}
	}
	return nil
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
