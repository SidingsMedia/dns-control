// SPDX-FileCopyrightText: 2025 Sidings Media
// SPDX-License-Identifier: MIT

package dnscontrol

import "errors"

var (
	ErrServerNotFound      = errors.New("server with provided id could not be found")
	ErrStatusNotOk         = errors.New("server returned an response code that was not 200 OK")
	ErrStructFieldNotFound = errors.New("attempted to lookup get name of struct field that doesn't exist")
)
