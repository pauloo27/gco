package config

import (
	"fmt"
	"os"
)

var defaultConfig = &Config{
	RemoveTrailingPeriods: true,
	PrefixPack:            "text",
}
var globalConfig *Config
var projectConfig *Config
var projectConfigNotFound bool

func GetGlobalConfig() (*Config, error) {
	if globalConfig == nil {
		userHome, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		globalConfig, err = LoadConfig(fmt.Sprintf("%s/.config/gommit.json", userHome))
		if err != nil {
			return nil, err
		}
	}
	return globalConfig, nil
}

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

func GetAnyConfig() (*Config, error) {
	projectConfig, err := GetProjectConfig()
	if !projectConfigNotFound {
		return projectConfig, err
	}
	return GetGlobalConfig()
}
