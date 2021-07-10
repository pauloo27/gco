package config

import (
	"encoding/json"
	"io"
	"os"
)

func LoadConfig(path string) (conf *Config, err error) {
	conf = &Config{}
	var file *os.File
	file, err = os.OpenFile(path, os.O_RDONLY, 0x770)
	if err != nil {
		return
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, conf)
	return
}
