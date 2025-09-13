// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package util

import (
	"log"
	"os"
	"strconv"
)

// Get the specified environment variable. If it doesn't exist, return
// the fallback instead
func SGetenv(key string, fallback string) string {
	val := os.Getenv(key)

	if len(val) == 0 {
		return fallback
	} else {
		return val
	}
}

// Get the specified environment variable as an integer. If is doesn't
// exist or cannot be converted to an int, return the fallback instead.
func IGetenv(key string, fallback int) int {
	val := os.Getenv(key)

	if len(val) == 0 {
		return fallback
	} else {
		res, err := strconv.Atoi(val)
		if err != nil {
			return fallback
		} else {
			return res
		}
	}
}

// Attempt to get the environment variable. If it is not set, log error
// and exit
func Mustgetenv(key string) string {
	val := os.Getenv(key)

	if len(val) == 0 {
		log.Fatalln("Failed to get", key, ". Environment variable not set")
	}
	return val
}
