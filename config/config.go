package config

import (
	"errors"
	"fmt"
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
	fileSyncPath   = filepath.Join(homeDir, configDir, "sync")
)

type Config struct {
	Settings Settings `toml:"settings"`
}

type Settings struct {
	ProfileName string `toml:"profile-name"`
	QueueUrl    string `toml:"queue-url"`
}

func (c *Config) ProfileName() string {
	return c.Settings.ProfileName
}

func (c *Config) QueueUrl() string {
	return c.Settings.QueueUrl
}

func GetConfig(profileName, queueUrl string) (*Config, error) {
	var cfg *Config

	// Check if all params provided
	if len(profileName) > 0 && len(queueUrl) > 0 {
		cfg = &Config{
			Settings{
				ProfileName: profileName,
				QueueUrl:    queueUrl,
			},
		}

		return cfg, nil
	} else {
    _, err := os.Stat(configFilePath)
    if os.IsNotExist(err) {
      writeDefaultConfig()
      e := fmt.Sprintf(
        "Fresh config created at %s. Please update config or pass in needed args",
        configFilePath,
      )
      return nil, errors.New(e)
    }

    cfg = readConfig()

    if profileName != "" {
      cfg.Settings.ProfileName = profileName
    }
    if queueUrl != "" {
      cfg.Settings.QueueUrl = queueUrl
    }


    return cfg, nil
	}
}

func (c *Config) EnsureConfigValues() error {
	if c.Settings.ProfileName == "" || c.Settings.QueueUrl == "" {
		e := fmt.Sprintf(
			"Please update config at %s or pass in needed args",
			configFilePath,
		)
		return errors.New(e)
	}

	return nil
}

func writeDefaultConfig() {
	c, err := toml.Marshal(Config{Settings: Settings{}})
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

	os.MkdirAll(fileSyncPath, 0750)
}

func readConfig() *Config {
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	var cfg Config
	err = toml.Unmarshal([]byte(data), &cfg)
	if err != nil {
		log.Fatal(err)
	}

	return &cfg
}
