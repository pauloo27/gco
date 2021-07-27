package config

type HookAction struct {
	Command         string
	Ask, ExitOnFail bool
}

type Config struct {
	PrefixPack       string
	AskFilesToCommit bool
	PreCommit        []HookAction
	PostCommit       []HookAction
	AskCISkip        bool
}
