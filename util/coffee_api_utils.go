package util

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func extractId(c *gin.Context) uint {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		panic(err)
	}

	return uint(id)
}
