// SPDX-FileCopyrightText: 2025 Sidings Media
// SPDX-License-Identifier: MIT

package domain

type TechnetiumResponse struct {
	Status            string `json:"status"`
	ErrorMessage      string `json:"errorMessage"`
	StackTrace        string `json:"stackTrace"`
	InnerErrorMessage string `json:"innerErrorMessage"`
}
