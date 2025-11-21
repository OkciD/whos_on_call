package config

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	ListenAddr string `mapstructure:"listenAddr"`
}

func ReadConfig(configFilePath string) (*Config, error) {
	fileName := strings.TrimSuffix(filepath.Base(configFilePath), filepath.Ext(configFilePath))
	dir := filepath.Dir(configFilePath)

	viper.AddConfigPath(dir)
	viper.SetConfigType("json")
	viper.SetConfigName(fileName)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var config Config

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}
