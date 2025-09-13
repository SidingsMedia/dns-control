// SPDX-FileCopyrightText: 2025 Sidings Media
// SPDX-License-Identifier: MIT

package model

type CachedResult struct {
	Id    string         `json:"id"`
	RData map[string]any `json:"data"`
	Ttl   string         `json:"ttl"`
}

type CacheEntry struct {
	Name         string         `json:"name"`
	Type         string         `json:"type"`
	CachedResult []CachedResult `json:"cached-results"`
}

type CacheResponse struct {
	AvailableDomains []string     `json:"domains"`
	Entries          []CacheEntry `json:"entries"`
}
