package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

// Config has config for moesia
type Config struct {
	From          string   `json:"from"`
	To            []string `json:"to"`
	Cc            []string `json:"cc"`
	GmailUserName string   `json:"gmailUserName"`
	GmailPassword string   `json:"gmailPassword"`
}

var filename = fmt.Sprintf("%s/.config/moesia/config.json", os.Getenv("HOME"))

// New returns a new instance of Config
func New() (config *Config, err error) {
	if _, err = os.Stat(filename); err != nil {
		if config, err = makeInitialConfigFile(filename); err != nil {
			err = fmt.Errorf("failed to make initial config file: %v", err)
			return
		}
	} else {
		if config, err = loadConfig(filename); err != nil {
			err = fmt.Errorf("failed to load config file: %v", err)
			return
		}
	}
	return
}

func makeInitialConfigFile(filename string) (config *Config, err error) {
	if err = os.MkdirAll(path.Dir(filename), 0700); err != nil {
		err = fmt.Errorf("failed to mkdir: %v", err)
		return
	}
	var file *os.File
	if file, err = os.Create(filename); err != nil {
		err = fmt.Errorf("failed to create %s: %v", filename, err)
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	if err = encoder.Encode(config); err != nil {
		err = fmt.Errorf("failed to encode Config %v: %v", config, err)
		return
	}
	return
}

func loadConfig(filename string) (config *Config, err error) {
	var file *os.File
	if file, err = os.Open(filename); err != nil {
		err = fmt.Errorf("failed to open %s: %v", filename, err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&config); err != nil {
		err = fmt.Errorf("failed to decode config: %v", err)
		return
	}
	return
}
