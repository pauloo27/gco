package config

import (
	"os"

	"github.com/Pauloo27/gco/utils/git"
)

var defaultConfig = &Config{
	PrefixPack: "text",
}

var globalConfig *Config
var projectConfig *Config
var projectConfigNotFound bool

var rcFileName = ".gcorc.json"

func StoreProjectConfig(conf *Config) error {
	gitRoot, err := git.GetRepositoryRoot()
	if err != nil {
		return err
	}
	return StoreConfig(conf, gitRoot+"/"+rcFileName)
}

func GetProjectConfig() (*Config, error) {
	if projectConfig == nil && !projectConfigNotFound {
		var err error
		gitRoot, err := git.GetRepositoryRoot()
		if err != nil {
			return nil, err
		}
		projectConfig, err = LoadConfig(gitRoot + "/" + rcFileName)
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
