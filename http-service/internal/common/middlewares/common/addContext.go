package common

import (
	"context"

	"github.com/gin-gonic/gin"
)

func AddContext(c *gin.Context) {
	c.Set("ctx", context.Background())
	c.Next()
}
