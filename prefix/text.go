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
			Key:         "fmt",
			Value:       "fmt",
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
			Key:         "del",
			Value:       "del",
		},
		{
			Description: "move code/files",
			Key:         "mov",
			Value:       "mov",
		},
		{
			Description: "rename code/files",
			Key:         "ren",
			Value:       "ren",
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
			Key:         "tag",
			Value:       "tag",
		},
		{
			Description: "fix typo",
			Key:         "typo",
			Value:       "typo",
		},
		{
			Description: "merge branches/pull requests",
			Key:         "merg",
			Value:       "merg",
		},
		{
			Description: "improve ui/ux/accessibility",
			Key:         "ux",
			Value:       "ux",
		},
		{
			Description: "refactor",
			Key:         "ref",
			Value:       "ref",
		},
		{
			Description: "internationalization",
			Key:         "i18n",
			Value:       "i18n",
		},
		{
			Description: "add/update .gitignore",
			Key:         "ig",
			Value:       "ig",
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
