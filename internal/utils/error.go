package utils

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidationErrorResponse(c *gin.Context, err error) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		errorsMap := map[string]string{}
		for _, fe := range ve {
			field := strings.ToLower(fe.Field())
			errorsMap[field] = fe.Tag() + " validation failed"
		}
		c.JSON(400, gin.H{
			"message": "Validation failed",
			"errors":  errorsMap,
		})
	}
}
