package m_json

import (
	"Spy-Cat-Agency/src/internal/missions/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JsonUpdateMissionValidator() gin.HandlerFunc {
	return func(c *gin.Context) {

		var completeReq dtos.CompletionStateRequest

		if err := c.ShouldBindJSON(&completeReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid completion state data"})
			c.Abort()
			return
		}

		c.Set("completeReq", completeReq)
		c.Next()
	}
}
