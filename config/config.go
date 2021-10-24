package config

type HookAction struct {
	Command         string
	Ask, ExitOnFail bool
}

type Config struct {
	Version          string
	PrefixPack       string
	AskFilesToCommit bool
	PreCommit        []HookAction
	PostCommit       []HookAction
	Modules          []string
	AskCISkip        bool
}
