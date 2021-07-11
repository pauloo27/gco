package prefix

var TextPrefix = PrefixPack{
	Name:        "text",
	Separator:   ": ",
	Description: "Text prefixes",
	Prefixes: []*Prefix{
		{
			Name:  "add feature",
			Value: "feat",
		},
		{
			Name:  "bug fix",
			Value: "fix",
		},
		{
			Name:  "add/update ci",
			Value: "ci",
		},
		{
			Name:  "reformat code",
			Value: "format",
		},
		{
			Name:  "add/update documentation",
			Value: "doc",
		},
		{
			Name:  "move or rename code/files",
			Value: "move",
		},
		{
			Name:  "add/update test",
			Value: "test",
		},
		{
			Name:  "upgrade dependency",
			Value: "up",
		},
		{
			Name:  "downgrade dependency",
			Value: "dep",
		},
		{
			Name:  "add/remove dependency",
			Value: "dep",
		},
		{
			Name:  "add tag/release",
			Value: "tag",
		},
		{
			Name:  "fix typo",
			Value: "typo",
		},
		{
			Name:  "merge branches/pull requests",
			Value: "merge",
		},
		{
			Name:  "improve ux/accessibility",
			Value: "ux",
		},
		{
			Name:  "improve code",
			Value: "improve",
		},
		{
			Name:  "refactor",
			Value: "refact",
		},
		{
			Name:  "internationalization",
			Value: "i18n",
		},
		{
			Name:  "work in progress",
			Value: "wip",
		},
	},
}
