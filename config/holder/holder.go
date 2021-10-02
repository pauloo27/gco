package holder

import (
	"os"
	"strings"

	"github.com/Pauloo27/gco/config"
	"github.com/Pauloo27/gco/utils/git"
)

var globalConfig = &config.Config{}
var projectConfig = &config.Config{}

var configFileName = "gco.json"

func StoreGlobalConfig(conf *config.Config) error {
	// tries to read env... if it's not found, use ~/.config
	configHome := os.Getenv("XDG_CONFIG_HOME")
	if configHome == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		configHome = home + "/.config/"
	}
	// add missing trailing `/` when it's needed
	if !strings.HasSuffix(configHome, "/") {
		configHome += "/"
	}
	return config.StoreConfig(conf, configHome+configFileName)
}

func StoreProjectConfig(conf *config.Config) error {
	gitRoot, err := git.GetRepositoryRoot()
	if err != nil {
		return err
	}
	return config.StoreConfig(conf, gitRoot+"/."+configFileName)
}

func GetGlobalConfig() (*config.Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	err = loadConfig(home+"/.config/"+configFileName, globalConfig)
	return globalConfig, err
}

func GetProjectConfig() (*config.Config, error) {
	gitRoot, err := git.GetRepositoryRoot()
	if err != nil {
		return nil, err
	}
	err = loadConfig(gitRoot+"/."+configFileName, projectConfig)
	return projectConfig, err
}

func GetProjectConfigOrGlobal() (conf *config.Config, isFromProject bool, err error) {
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
