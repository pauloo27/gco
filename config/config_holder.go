package config

import (
	"os"
)

var defaultConfig = &Config{
	PrefixPack: "text",
}

var globalConfig *Config
var projectConfig *Config
var projectConfigNotFound bool

func GetProjectConfig() (*Config, error) {
	if projectConfig == nil && !projectConfigNotFound {
		var err error
		// TODO: use git root
		projectConfig, err = LoadConfig("./.gommitrc.json")
		if err != nil {
			if os.IsNotExist(err) {
				projectConfigNotFound = true
				return nil, err
			} else {
				return nil, err
			}
		}
	}
	return projectConfig, nil
}
