package utils

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func ParseDateParam(c *gin.Context, key string) (*time.Time, error) {
	value := c.Query(key)
	if value == "" {
		return nil, nil
	}

	t, err := time.Parse("2006-01-02", value)
	if err != nil {
		return nil, fmt.Errorf("invalid %s format, use YYYY-MM-DD", key)
	}

	return &t, nil
}
