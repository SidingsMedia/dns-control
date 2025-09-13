// SPDX-FileCopyrightText: 2023 Sidings Media
// SPDX-License-Identifier: MIT

package util

import (
	"net/http"
	"strings"

	"github.com/SidingsMedia/dns-control/dnscontrol/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Send a standard bad request response but include a list of fields
// that suffered from binding errors
func SendBadRequestFieldNames(ctx *gin.Context, validationError validator.ValidationErrors) {
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
