package util

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type ModeConfig struct {
	Mode string `json:"mode"`
}

type Config struct {
	Template string  `json:"template"`
	Compile  *string `json:"compile"`
	Run      string  `json:"run"`

	Bundle      *string           `json:"bundle"`
	BundleACL   *string           `json:"bundle_acl"`
	RandomTest  RandomTestConfig  `json:"random_test"`
	Interactive InteractiveConfig `json:"interactive"`
}

//TODO Integrate random test and interactive commands internally.

type RandomTestConfig struct {
	Naive     string  `json:"naive"`
	Generator string  `json:"generator"`
	Compile   *string `json:"compile"`
	Run       string  `json:"run"`
}

type InteractiveConfig struct {
	Judge   string  `json:"judge"`
	Compile *string `json:"compile"`
	Run     string  `json:"run"`
}

func GetMode() (string, error) {
	configPath := ResolveConfigPath("config.json")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return "", fmt.Errorf("Config file not found at %s", configPath)
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return "", fmt.Errorf("Failed to read config file: %w", err)
	}

	var lang ModeConfig
	if err := json.Unmarshal(data, &lang); err != nil {
		return "", fmt.Errorf("Failed to unmarshal config: %w", err)
	}

	return lang.Mode, nil
}

func LoadConfig() (Config, error) {
	lang, err := GetMode()
	if err != nil {
		return Config{}, err
	}
	configPath := ResolveConfigPath(filepath.Join("modes", lang+".json"))

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return Config{}, fmt.Errorf("Config file not found at %s", configPath)
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("Failed to read config file: %w", err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, fmt.Errorf("Failed to unmarshal config: %w", err)
	}

	return config, nil
}

func ResolveConfigPath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}

	homeDir := os.Getenv("HOME")
	return filepath.Join(homeDir, ".config", "ojx", path)
}
