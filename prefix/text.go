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
			Description: "add/update ci",
			Key:         "ci",
			Value:       "ci",
		},
		{
			Description: "reformat code",
			Key:         "format",
			Value:       "format",
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
			Description: "delete code",
			Key:         "delete",
			Value:       "delete",
		},
		{
			Description: "rename/move code/files",
			Key:         "rename",
			Value:       "rename",
		},
		{
			Description: "add/update test",
			Key:         "test",
			Value:       "test",
		},
		{
			Description: "add/remove/update dependency",
			Key:         "dep",
			Value:       "dep",
		},
		{
			Description: "add tag/release",
			Key:         "release",
			Value:       "release",
		},
		{
			Description: "fix typo",
			Key:         "typo",
			Value:       "typo",
		},
		{
			Description: "merge branches/pull requests",
			Key:         "merge",
			Value:       "merge",
		},
		{
			Description: "improve ui/ux/accessibility",
			Key:         "ux",
			Value:       "ux",
		},
		{
			Description: "refactor",
			Key:         "refact",
			Value:       "refact",
		},
		{
			Description: "internationalization",
			Key:         "i18n",
			Value:       "i18n",
		},
		{
			Description: "add/update .gitignore",
			Key:         "ignore",
			Value:       "ignore",
		},
		{
			Description: "add/remove development related stuff",
			Key:         "dev",
			Value:       "dev",
		},
		{
			Description: "work in progress",
			Key:         "wip",
			Value:       "wip",
		},
	},
}
