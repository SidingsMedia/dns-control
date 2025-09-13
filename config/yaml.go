// SPDX-FileCopyrightText: 2025 Sidings Media
// SPDX-License-Identifier: MIT

package config

type Server struct {
	Target string `yaml:"target"`
	Name   string `yaml:"name"`
	Token  string `yaml:"token"`
}

type ConfigFile struct {
	Servers        []Server `yaml:"servers"`
	BindAddr       string   `yaml:"bind"`
	TrustedProxies []string `yaml:"trusted-proxies"`
	Debug          bool     `yaml:"debug"`
}
