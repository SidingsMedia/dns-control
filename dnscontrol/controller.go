// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package dnscontrol

import (
	"net/http"
	"strings"

	"github.com/SidingsMedia/dns-control/dnscontrol/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Controller interface {
	Health(ctx *gin.Context)
	ListServers(ctx *gin.Context)
}

type controller struct {
	service Service
}

// Format and send the JSON response. If the indent query parameter is
// present, this will pretty print the JSON for the client.
func formatJson(ctx *gin.Context, code int, obj any) {
	_, exists := ctx.GetQuery("indent")
	if exists {
		ctx.IndentedJSON(code, obj)
	} else {
		ctx.JSON(code, obj)
	}
}

// Send a standard bad request response but include a list of fields
// that suffered from binding errors
func sendBadRequestFieldNames(ctx *gin.Context, validationError validator.ValidationErrors) {
	response := &model.BadRequest{
		Code:    http.StatusBadRequest,
		Message: "Your request is malformed",
	}
	// Itterate through errors and add field name and condition to fields
	for _, malformedField := range validationError {
		field := malformedField.Field()
		response.Fields = append(response.Fields, model.Fields{
			Field:     strings.ToLower(field[:1]) + field[1:],
			Condition: malformedField.ActualTag(),
		})
	}
	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
}

func (controller controller) HealthCheck(ctx *gin.Context) {
	// if err := controller.service.HealthCheck(); err != nil {
	//     ctx.String(http.StatusServiceUnavailable, "unhealthy")
	//     ctx.Abort()
	//     return
	// }
	ctx.String(http.StatusOK, "healthy")
}

func (controller controller) ListServers(ctx *gin.Context) {
	formatJson(ctx, http.StatusOK, controller.service.ListServers())
}

func NewController(engine *gin.Engine, Service Service) {
	controller := &controller{
		service: Service,
	}
	api := engine.Group("")
	{
		api.GET("health", controller.HealthCheck)
		api.GET("servers", controller.ListServers)
	}
}
