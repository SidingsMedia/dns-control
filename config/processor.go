// SPDX-FileCopyrightText: 2025 Sidings Media
// SPDX-License-Identifier: MIT

package config

import (
	"log/slog"
	"os"

	yaml "github.com/goccy/go-yaml"
	"github.com/jinzhu/copier"
)

func populateDefaults(config *ConfigFile) {
	if config.TrustedProxies == nil {
		config.TrustedProxies = DefaultTrustedProxies
	}

	if config.BindAddr == "" {
		config.BindAddr = DefaultBindAddr
	}
}

func logSanitized(config ConfigFile) {
	for i := range config.Servers {
		config.Servers[i].Token = "***"
	}
	slog.Info("Read config file", "config", config)
}

func ReadConfigFile(path string) (*ConfigFile, error) {
	slog.Info("Reading configuration file", "path", path)
	file, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var config ConfigFile
	if err := yaml.Unmarshal(file, &config); err != nil {
		return nil, err
	}

	populateDefaults(&config)
	var sanitized ConfigFile
	copier.CopyWithOption(&sanitized, config, copier.Option{DeepCopy: true})
	logSanitized(sanitized)

	return &config, nil
}
