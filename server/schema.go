// SPDX-FileCopyrightText: 2025 Sidings Media
// SPDX-License-Identifier: MIT

package server

type GetCacheRequest struct {
	Domain  string   `form:"domain"`
	Servers []string `form:"server" binding:"required"`
}
