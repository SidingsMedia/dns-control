// SPDX-FileCopyrightText: 2025 Sidings Media
// SPDX-License-Identifier: MIT

package dnscontrol

import (
	"errors"
	"log/slog"
	"net/http"
	"reflect"
	"strings"

	"github.com/SidingsMedia/dns-control/dnscontrol/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Controller interface {
	Health(ctx *gin.Context)
	ListServers(ctx *gin.Context)
	GetCache(ctx *gin.Context)
	DeleteCacheEntry(ctx *gin.Context)
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

// Attempt to get the name of the field. Will attempt to use the form
// tag for GET request and the json tag for all other requests, falling
// back to a lowercase of the field name as required.
func getFieldName(ctx *gin.Context, malformedField validator.FieldError, typ reflect.Type) (string, error) {
	field, ok := typ.FieldByName(malformedField.StructField())
	if !ok {
		return "", ErrStructFieldNotFound
	}

	tag := ""
	if ctx.Request.Method == "GET" {
		// Probably query data
		tag = field.Tag.Get("form")
	} else {
		tag = field.Tag.Get("json")
	}

	if tag == "" {
		tag = strings.ToLower(field.Name[:1]) + field.Name[1:]
	}

	return tag, nil
}

// Send a standard bad request response but include a list of fields
// that suffered from binding errors
func sendBadRequestFieldNames(ctx *gin.Context, validationError validator.ValidationErrors, typ reflect.Type) {
	response := &model.BadRequest{
		GeneralError: model.GeneralError{
			Code:    http.StatusBadRequest,
			Message: "Your request is malformed",
		},
	}

	// Iterate through errors and add field name and condition to fields
	for _, malformedField := range validationError {
		field, err := getFieldName(ctx, malformedField, typ)

		if err != nil {
			slog.Error("Failed to get field name when processing bad request.", "error", err)
			formatJson(ctx, http.StatusInternalServerError,
				model.GeneralError{
					Code:    http.StatusInternalServerError,
					Message: "An internal error occurred",
				},
			)
		}

		response.Fields = append(response.Fields, model.Fields{
			Field:     field,
			Condition: malformedField.ActualTag(),
		})
	}

	formatJson(ctx, http.StatusBadRequest, response)
	ctx.Abort()
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

func (controller controller) GetCache(ctx *gin.Context) {
	queryParams := GetCacheRequest{}

	if err := ctx.BindQuery(&queryParams); err != nil && errors.As(err, &validator.ValidationErrors{}) {
		sendBadRequestFieldNames(ctx, err.(validator.ValidationErrors), reflect.TypeOf(GetCacheRequest{}))
		return
	} else if err != nil {
		formatJson(ctx, http.StatusBadRequest, model.GeneralError{
			Code:    http.StatusBadRequest,
			Message: "Request was malformed",
		})

		ctx.Abort()
		return
	}

	response, err := controller.service.GetCache(queryParams.Domain, queryParams.Servers)
	if err != nil {
		switch err {
		case ErrServerNotFound:
			formatJson(ctx, http.StatusNotFound, model.GeneralError{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			})
			ctx.Abort()
			return
		}

		formatJson(ctx, http.StatusInternalServerError, model.GeneralError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		ctx.Abort()
		return
	}

	formatJson(ctx, http.StatusOK, response)
}

func (controller controller) DeleteCacheEntry(ctx *gin.Context) {
	queryParams := GetCacheRequest{}

	if err := ctx.BindQuery(&queryParams); err != nil && errors.As(err, &validator.ValidationErrors{}) {
		sendBadRequestFieldNames(ctx, err.(validator.ValidationErrors), reflect.TypeOf(GetCacheRequest{}))
		return
	} else if err != nil {
		formatJson(ctx, http.StatusBadRequest, model.GeneralError{
			Code:    http.StatusBadRequest,
			Message: "Request was malformed",
		})

		ctx.Abort()
		return
	}

	response, err := controller.service.DeleteCacheEntry(queryParams.Domain, queryParams.Servers)
	if err != nil {
		switch err {
		case ErrServerNotFound:
			formatJson(ctx, http.StatusNotFound, model.GeneralError{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			})
			ctx.Abort()
			return
		}

		formatJson(ctx, http.StatusInternalServerError, model.GeneralError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		ctx.Abort()
		return
	}

	if response != nil {
		formatJson(ctx, response.Code, response)
		ctx.Abort()
	}

	ctx.Status(http.StatusNoContent)
}

func NewController(engine *gin.Engine, Service Service) {
	controller := &controller{
		service: Service,
	}
	api := engine.Group("")
	{
		api.GET("health", controller.HealthCheck)
		api.GET("servers", controller.ListServers)
		api.GET("cache", controller.GetCache)
		api.DELETE("cache", controller.DeleteCacheEntry)
	}
}
