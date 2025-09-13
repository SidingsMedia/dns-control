// SPDX-FileCopyrightText: 2023-2025 Sidings Media
// SPDX-License-Identifier: MIT

package model

type GeneralError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type AffectedServer struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

type PerServerFail struct {
	GeneralError
	AffectedServers []AffectedServer `json:"servers"`
}

type Fields struct {
	Field     string `json:"field"`
	Condition string `json:"condition"`
}

// Standardised error response schema
type BadRequest struct {
	GeneralError
	Fields []Fields `json:"fields"`
}
