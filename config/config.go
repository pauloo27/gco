package config

type Prefix struct {
	ID, Name, Value string
}

type Config struct {
	RemoveTrailingPeriods bool
	PrefixPack            string
	Prefixes              []Prefix
}
