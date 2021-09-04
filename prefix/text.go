package prefix

var TextPrefix = PrefixPack{
	Name:        "text",
	Separator:   ": ",
	Description: "Text prefixes",
	Prefixes: []*Prefix{
		{
			Description: "add feature",
			Name:        "feat",
		},
		{
			Description: "bug fix",
			Name:        "fix",
		},
		{
			Description: "add/update ci",
			Name:        "ci",
		},
		{
			Description: "reformat code",
			Name:        "format",
		},
		{
			Description: "add/update documentation",
			Name:        "doc",
		},
		{
			Description: "initial commit",
			Name:        "init",
		},
		{
			Description: "delete code",
			Name:        "delete",
		},
		{
			Description: "move or rename code/files",
			Name:        "move",
		},
		{
			Description: "add/update test",
			Name:        "test",
		},
		{
			Description: "add/remove/update dependency",
			Name:        "dep",
		},
		{
			Description: "add tag/release",
			Name:        "release",
		},
		{
			Description: "fix typo",
			Name:        "typo",
		},
		{
			Description: "merge branches/pull requests",
			Name:        "merge",
		},
		{
			Description: "improve ux/accessibility",
			Name:        "ux",
		},
		{
			Description: "improve code",
			Name:        "improve",
		},
		{
			Description: "refactor",
			Name:        "refact",
		},
		{
			Description: "internationalization",
			Name:        "i18n",
		},
		{
			Description: "add/update .gitignore",
			Name:        "ignore",
		},
		{
			Description: "add/remove development related stuff",
			Name:        "dev",
		},
		{
			Description: "work in progress",
			Name:        "wip",
		},
	},
}
