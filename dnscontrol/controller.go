// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package dnscontrol

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	Health(ctx *gin.Context)
}

type controller struct {
	service Service
}

func (controller controller) HealthCheck(ctx *gin.Context) {
	// if err := controller.service.HealthCheck(); err != nil {
	//     ctx.String(http.StatusServiceUnavailable, "unhealthy")
	//     ctx.Abort()
	//     return
	// }
	ctx.String(http.StatusOK, "healthy")
}

func NewController(engine *gin.Engine, Service Service) {
	controller := &controller{
		service: Service,
	}
	api := engine.Group("")
	{
		api.GET("health", controller.HealthCheck)
	}
}
