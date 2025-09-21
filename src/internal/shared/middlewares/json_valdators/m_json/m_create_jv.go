package m_json

import (
	"Spy-Cat-Agency/src/internal/missions/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JsonCreateMissionValidator() gin.HandlerFunc {
	return func(c *gin.Context) {

		var missionReq dtos.MissionCreateRequest

		if err := c.ShouldBindJSON(&missionReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mission data"})
			c.Abort()
			return
		}

		c.Set("missionReq", missionReq)
		c.Next()
	}
}
