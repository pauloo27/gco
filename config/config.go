package config

type Prefix struct {
	ID, Name, Value string
}

type Config struct {
	PrefixPack string
	Prefixes   []Prefix
}
