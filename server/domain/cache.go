// SPDX-FileCopyrightText: 2025 Sidings Media
// SPDX-License-Identifier: MIT

package domain

type CacheRecord struct {
	Name  string         `json:"name"`
	Type  string         `json:"type"`
	Ttl   string         `json:"ttl"`
	RData map[string]any `json:"rData"`
}

type CacheResult struct {
	TechnetiumResponse
	Response struct {
		Domain  string        `json:"domain"`
		Zones   []string      `json:"zones"`
		Records []CacheRecord `json:"records"`
	} `json:"response"`
}

type PerServerFail struct {
	Id  string
	Err error
}
