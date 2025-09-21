package id_validators

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SpyCatIdValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("spycatId")

		id, err := uuid.Parse(idStr)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid spycat Id"})
			c.Abort()
			return
		}

		c.Set("spycatId", id)
		c.Next()
	}
}
