// SPDX-FileCopyrightText: 2025 Sidings Media
// SPDX-License-Identifier: MIT

package model

type List[T any] struct {
	Results []T `json:"results"`
}
