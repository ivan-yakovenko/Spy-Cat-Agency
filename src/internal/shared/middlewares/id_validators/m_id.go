package id_validators

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func MissionIdValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("missionId")

		id, err := uuid.Parse(idStr)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mission Id"})
			c.Abort()
			return
		}

		c.Set("missionId", id)
		c.Next()
	}
}
