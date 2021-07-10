package config

import (
	"encoding/json"
	"os"
)

func StoreConfig(c *Config, path string) error {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	if _, err := os.Stat(path); err == nil {
		return os.ErrExist
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	return err
}
