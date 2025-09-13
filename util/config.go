// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package util

// Runtime config
var (
	// Address to bind server to
	BindAddr string
	// Proxies from which to trust alternative client IP headers
	TrustedProxies []string
)

// Environment variables
const (
	BindAddrEnv       = "BIND_ADDRESS"
	TrustedProxiesEnv = "TRUSTED_PROXIES"
)

// Defaults
const (
	DefaultBindAddr       = "[::]:3000"
	DefaultTrustedProxies = "*"
)
