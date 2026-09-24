package util

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Program     ProgramConfig     `json:"program"`
	Template    string            `json:"template"`
	RandomTest  RandomTestConfig  `json:"random_test"`
	Interactive InteractiveConfig `json:"interactive"`
}

type ProgramConfig struct {
	Compile   *string `json:"compile"`
	Run       string  `json:"run"`
	Bundle    *string `json:"bundle"`
	BundleACL *string `json:"bundle_acl"`
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

func LoadConfig() (Config, error) {
	configPath := ResolveConfigPath("config.json")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return Config{}, fmt.Errorf("Config file not found at %s", configPath)
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("Failed to read config file: %v", err)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, fmt.Errorf("Failed to unmarshal config: %v", err)
	}

	return config, nil
}

func ResolveConfigPath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}

	homeDir := os.Getenv("HOME")
	return filepath.Join(homeDir, ".config", "atc", path)
}
