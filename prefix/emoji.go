package prefix

func init() {
	Packs = append(Packs, EmojiPack)
}

var EmojiPack = PrefixPack{
	Name:        "emoji",
	Separator:   " ",
	Description: "Emoji prefixes",
	Prefixes: []*Prefix{
		{
			Description: "add feature",
			Key:         "feat",
			Value:       "✨",
		},
		{
			Description: "bug fix",
			Key:         "fix",
			Value:       "🐛",
		},
		{
			Description: "add/update ci",
			Key:         "ci",
			Value:       "👷",
		},
		{
			Description: "reformat code",
			Key:         "format",
			Value:       "📏",
		},
		{
			Description: "add/update documentation",
			Key:         "doc",
			Value:       "📖",
		},
		{
			Description: "initial commit",
			Key:         "init",
			Value:       "🎉",
		},
		{
			Description: "delete code",
			Key:         "delete",
			Value:       "🔥",
		},
		{
			Description: "rename/move code/files",
			Key:         "rename",
			Value:       "🚚",
		},
		{
			Description: "add/update test",
			Key:         "test",
			Value:       "✅",
		},
		{
			Description: "add/remove/update dependency",
			Key:         "dep",
			Value:       "🔩",
		},
		{
			Description: "add tag/release",
			Key:         "release",
			Value:       "🔖",
		},
		{
			Description: "fix typo",
			Key:         "typo",
			Value:       "✏️",
		},
		{
			Description: "merge branches/pull requests",
			Key:         "merge",
			Value:       "🔀",
		},
		{
			Description: "improve ui/ux/accessibility",
			Key:         "ux",
			Value:       "ux",
		},
		{
			Description: "refactor",
			Key:         "refact",
			Value:       "♻️",
		},
		{
			Description: "internationalization",
			Key:         "i18n",
			Value:       "🌐",
		},
		{
			Description: "add/update .gitignore",
			Key:         "ignore",
			Value:       "🙈",
		},
		{
			Description: "add/remove development related stuff",
			Key:         "dev",
			Value:       "💻",
		},
		{
			Description: "work in progress",
			Key:         "wip",
			Value:       "🚧",
		},
	},
}
