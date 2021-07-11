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
			Description: "move or rename code/files",
			Name:        "move",
		},
		{
			Description: "add/update test",
			Name:        "test",
		},
		{
			Description: "upgrade dependency",
			Name:        "up",
		},
		{
			Description: "downgrade dependency",
			Name:        "dep",
		},
		{
			Description: "add/remove dependency",
			Name:        "dep",
		},
		{
			Description: "add tag/release",
			Name:        "tag",
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
			Description: "work in progress",
			Name:        "wip",
		},
	},
}
