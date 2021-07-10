package prefix

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
