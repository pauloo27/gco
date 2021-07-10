package config

import (
	"fmt"
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
		projectConfig, err = LoadConfig("./.gommit.json")
		if err != nil {
			if os.IsNotExist(err) {
				projectConfigNotFound = true
			} else {
				return nil, err
			}
		}
	}
	return projectConfig, nil
}
