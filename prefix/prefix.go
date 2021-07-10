package prefix

type Prefix struct {
	Name, Value string
}

type PrefixPack struct {
	Name, Description string
	Prefixes          map[string]*Prefix
}
