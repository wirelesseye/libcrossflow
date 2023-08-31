package config

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"gopkg.in/yaml.v3"
)

type Config struct {
	internel configInternel
}

type ShareSpace struct {
	Files map[string]string
}

type configInternel struct {
	ShareSpaces map[string]ShareSpace
}

func defaultConfig() Config {
	return Config{
		internel: configInternel{
			ShareSpaces: map[string]ShareSpace{"default": {
				Files: map[string]string{"home": "~"},
			}},
		},
	}
}

var lock = &sync.Mutex{}
var configInstance *Config

func Initialize() {
	if configInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if configInstance == nil {
			dirname, _ := os.UserConfigDir()
			configFilePath := filepath.Join(dirname, "crossflow", "config.yaml")
			log.Printf("Loading config from %s\n", configFilePath)
			dat, err := os.ReadFile(configFilePath)
			if os.IsNotExist(err) {
				c := defaultConfig()
				data, _ := yaml.Marshal(c.internel)
				os.MkdirAll(filepath.Join(dirname, "crossflow"), 0755)
				os.WriteFile(configFilePath, data, 0644)
				configInstance = &c
			} else {
				internel := configInternel{}
				yaml.Unmarshal(dat, &internel)
				configInstance = &Config{internel}
			}
		}
	}
}

func GetConfig() *Config {
	if configInstance == nil {
		Initialize()
	}
	return configInstance
}

func (config *Config) GetRawData() *configInternel {
	return &config.internel
}
