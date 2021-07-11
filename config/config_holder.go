package config

import (
	"os"

	"github.com/Pauloo27/gommit/utils"
)

var defaultConfig = &Config{
	PrefixPack: "text",
}

var globalConfig *Config
var projectConfig *Config
var projectConfigNotFound bool

func StoreProjectConfig(conf *Config) error {
	gitRoot, err := utils.GetRepositoryRoot()
	if err != nil {
		return err
	}
	return StoreConfig(conf, gitRoot+"/.gotmmitrc.json")
}

func GetProjectConfig() (*Config, error) {
	if projectConfig == nil && !projectConfigNotFound {
		var err error
		gitRoot, err := utils.GetRepositoryRoot()
		if err != nil {
			return nil, err
		}
		projectConfig, err = LoadConfig(gitRoot + "/.gommitrc.json")
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
