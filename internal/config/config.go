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
)

type Config struct {
	Settings Settings `toml:"settings"`
}

type Settings struct {
	ProfileName string `toml:"profile-name"`
	SyncBucket  string `toml:"sync-bucket-name"`
	QueueUrl    string `toml:"queue-url"`
}

func (c *Config) ProfileName() string {
	return c.Settings.ProfileName
}

func (c *Config) SyncBucket() string {
	return c.Settings.SyncBucket
}

func (c *Config) QueueUrl() string {
	return c.Settings.QueueUrl
}

func EnsureConfig(profileName, syncBucket, queueUrl string) (*Config, error) {
	_, err := os.Stat(configFilePath)
	if os.IsNotExist(err) {
		writeDefaultConfig()
		e := fmt.Sprintf(
			"Fresh config created at %s. Please update config or pass in needed args",
			configFilePath,
		)
		return nil, errors.New(e)
	}

	cfg := readConfig()

	if profileName != "" {
		cfg.Settings.ProfileName = profileName
	}
	if syncBucket != "" {
		cfg.Settings.SyncBucket = syncBucket
	}
	if queueUrl != "" {
		cfg.Settings.QueueUrl = queueUrl
	}

	err = cfg.ensureConfigValues()
	if err != nil {
		return nil, err 
	}

	return cfg, nil
}

func (c *Config) ensureConfigValues() error {
	if c.Settings.ProfileName == "" || c.Settings.SyncBucket == "" || c.Settings.QueueUrl == "" {
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
