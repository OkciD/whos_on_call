package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func ReadConfig[C any](configFilePath string) (*C, error) {
	absConfigFilePath, err := filepath.Abs(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path to config: %w", err)
	}

	configFileContents, err := os.ReadFile(absConfigFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %s: %w", absConfigFilePath, err)
	}

	var config C
	err = json.Unmarshal(configFileContents, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file %s: %w", absConfigFilePath, err)
	}

	return &config, nil
}
