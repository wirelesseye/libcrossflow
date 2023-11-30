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

type ShareSpaceConfig struct {
	Files map[string]string
}

type configInternel struct {
	ShareSpaces map[string]ShareSpaceConfig
}

func defaultConfig() Config {
	homeDir, _ := os.UserHomeDir()

	return Config{
		internel: configInternel{
			ShareSpaces: map[string]ShareSpaceConfig{"Default": {
				Files: map[string]string{"home": homeDir},
			}},
		},
	}
}

var lock = &sync.Mutex{}
var configInstance *Config

func LoadConfig() {
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
		LoadConfig()
	}
	return configInstance
}

func (config *Config) GetRawData() *configInternel {
	return &config.internel
}
