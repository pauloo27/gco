package holder

import (
	"os"

	"github.com/Pauloo27/gco/config"
	"github.com/Pauloo27/gco/utils/git"
)

var defaultConfig = &config.Config{
	PrefixPack: "text",
}

var globalConfig = &config.Config{}
var projectConfig = &config.Config{}

var rcFileName = ".gco.json"

func StoreGlobalConfig(conf *config.Config) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	return config.StoreConfig(conf, home+"/"+rcFileName)
}

func StoreProjectConfig(conf *config.Config) error {
	gitRoot, err := git.GetRepositoryRoot()
	if err != nil {
		return err
	}
	return config.StoreConfig(conf, gitRoot+"/"+rcFileName)
}

func GetGlobalConfig() (*config.Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	err = loadConfig(home+"/"+rcFileName, globalConfig)
	return globalConfig, err
}

func GetProjectConfig() (*config.Config, error) {
	gitRoot, err := git.GetRepositoryRoot()
	if err != nil {
		return nil, err
	}
	err = loadConfig(gitRoot+"/"+rcFileName, projectConfig)
	return projectConfig, err
}

func GetProjectConfigOrGlobal() (conf *config.Config, err error, isFromProject bool) {
	conf, err = GetProjectConfig()
	if err == nil {
		isFromProject = true
		return
	}

	if os.IsNotExist(err) {
		conf, err = GetGlobalConfig()
	}

	return
}

func loadConfig(path string, conf *config.Config) error {
	loadedConf, err := config.LoadConfig(path)
	if err != nil {
		return err
	}
	*conf = *loadedConf
	return nil
}
