package holder

import (
	"os"

	"github.com/Pauloo27/gco/config"
	"github.com/Pauloo27/gco/utils/git"
)

var defaultConfig = &config.Config{
	PrefixPack: "text",
}

var globalConfig *config.Config
var projectConfig *config.Config
var projectConfigNotFound bool

var rcFileName = ".gco.json"

func StoreProjectConfig(conf *config.Config) error {
	gitRoot, err := git.GetRepositoryRoot()
	if err != nil {
		return err
	}
	return config.StoreConfig(conf, gitRoot+"/"+rcFileName)
}

func GetProjectConfig() (*config.Config, error) {
	if projectConfig == nil && !projectConfigNotFound {
		var err error
		gitRoot, err := git.GetRepositoryRoot()
		if err != nil {
			return nil, err
		}
		projectConfig, err = config.LoadConfig(gitRoot + "/" + rcFileName)
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
