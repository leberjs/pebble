package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

var (
	homeDir, _     = os.UserHomeDir()
	configDir      = ".pebble"
	configDirPath  = filepath.Join(homeDir, configDir)
	configFilePath = filepath.Join(homeDir, configDir, "pebble.toml")
)

type PebbleConfig struct {
	Settings Settings `toml:"settings"`
}

type Settings struct {
	AwsProfile string `toml:"profile-name"`
	SyncBucket string `toml:"sync-bucket-name"`
	QueueUrl   string `toml:"queue-url"`
}

func EnsureConfig() (bool, string) {
	_, err := os.Stat(configFilePath)
	if os.IsNotExist(err) {
		c, err := toml.Marshal(PebbleConfig{Settings: Settings{}})
		if err != nil {
			log.Fatal(err)
		}

		err = os.MkdirAll(configDirPath, 0750)
		if err != nil {
			log.Fatal(err)
		}

		err = os.WriteFile(configFilePath, []byte(c), 0660)
		if err != nil {
			log.Fatal(err)
		}

		return false, configFilePath
	}

	return true, ""
}

func ReadConfig() *PebbleConfig {
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	var cfg PebbleConfig
	err = toml.Unmarshal([]byte(data), &cfg)
	if err != nil {
		log.Fatal(err)
	}

	return &cfg
}

func (cfg *PebbleConfig) EnsureConfigValues() (bool, string) {
	if cfg.Settings.AwsProfile == "" || cfg.Settings.SyncBucket == "" || cfg.Settings.QueueUrl == "" {
		return false, configFilePath
	}

	return true, ""
}
