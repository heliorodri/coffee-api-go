package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ExtractId(c *gin.Context) uint {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		panic(err)
	}

	return uint(id)
}

func FormatId(id uint) string {
	return strconv.FormatUint(uint64(id), 10)
}
