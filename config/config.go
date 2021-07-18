package config

type HookAction struct {
	Command         string
	Ask, ExitOnFail bool
}

type Config struct {
	PrefixPack string
	PreCommit  []HookAction
	PostCommit []HookAction
	AskCISkip  bool
}
