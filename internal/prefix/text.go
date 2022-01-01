package prefix

func init() {
	Packs = append(Packs, TextPack)
}

var TextPack = PrefixPack{
	Name:        "text",
	Separator:   ": ",
	Description: "Text prefixes",
	Prefixes: []*Prefix{
		{
			Description: "add feature",
			Key:         "feat",
			Value:       "feat",
		},
		{
			Description: "bug fix",
			Key:         "fix",
			Value:       "fix",
		},
		{
			Description: "repair a small thing",
			Key:         "rep",
			Value:       "rep",
		},
		{
			Description: "merge something",
			Key:         "merge",
			Value:       "merge",
		},
		{
			Description: "revert a commit",
			Key:         "rev",
			Value:       "rev",
		},
		{
			Description: "add/update ci",
			Key:         "ci",
			Value:       "ci",
		},
		{
			Description: "style/reformat code",
			Key:         "style",
			Value:       "style",
		},
		{
			Description: "add/update documentation",
			Key:         "doc",
			Value:       "doc",
		},
		{
			Description: "initial commit",
			Key:         "init",
			Value:       "init",
		},
		{
			Description: "add/update test",
			Key:         "test",
			Value:       "test",
		},
		{
			Description: "improve performance",
			Key:         "perf",
			Value:       "perf",
		},
		{
			Description: "refactor",
			Key:         "ref",
			Value:       "ref",
		},
		{
			Description: "add/remove development related stuff",
			Key:         "dev",
			Value:       "dev",
		},
	},
}
