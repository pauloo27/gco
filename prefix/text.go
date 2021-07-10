package prefix

var TextPrefix = PrefixPack{
	Name:        "Text",
	Description: "Text prefixes",
	Prefixes: map[string]*Prefix{
		"feat": {
			Name:  "add feature",
			Value: "feat",
		},
		"fix": {
			Name:  "bug fix",
			Value: "fix",
		},
		"ci": {
			Name:  "add/update ci",
			Value: "ci",
		},
		"format": {
			Name:  "reformat code",
			Value: "format",
		},
		"doc": {
			Name:  "add/update documentation",
			Value: "doc",
		},
		"move": {
			Name:  "move or rename code/files",
			Value: "move",
		},
		"test": {
			Name:  "add/update test",
			Value: "test",
		},
		"up": {
			Name:  "upgrade dependency",
			Value: "up",
		},
		"down": {
			Name:  "downgrade dependency",
			Value: "dep",
		},
		"dep": {
			Name:  "add/remove dependency",
			Value: "dep",
		},
		"tag": {
			Name:  "add tag/release",
			Value: "tag",
		},
		"typo": {
			Name:  "fix typo",
			Value: "typo",
		},
		"merge": {
			Name:  "merge branches/pull requests",
			Value: "merge",
		},
		"ux": {
			Name:  "improve ux/accessibility",
			Value: "ux",
		},
		"refact": {
			Name:  "refactor",
			Value: "refact",
		},
		"i18n": {
			Name:  "internationalization",
			Value: "i18n",
		},
		"wip": {
			Name:  "work in progress",
			Value: "wip",
		},
	},
}
