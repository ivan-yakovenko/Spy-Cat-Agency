package t_json

import (
	"Spy-Cat-Agency/src/internal/missions/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JsonCreateTargetValidator() gin.HandlerFunc {
	return func(c *gin.Context) {

		var targetReq dtos.MissionTargetRequest

		if err := c.ShouldBindJSON(&targetReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid target data"})
			return
		}

		c.Set("targetReq", targetReq)
		c.Next()
	}
}
