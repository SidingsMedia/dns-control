// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package model

type Fields struct {
   Field     string `json:"field"`
   Condition string `json:"condition"`
}

// Standardised error response schema
type BadRequest struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
    Fields  []Fields `json:"fields"`
}
