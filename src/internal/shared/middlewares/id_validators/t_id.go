package id_validators

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TargetIdValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("targetId")

		id, err := uuid.Parse(idStr)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid target Id"})
			c.Abort()
			return
		}

		c.Set("targetId", id)
		c.Next()
	}
}
