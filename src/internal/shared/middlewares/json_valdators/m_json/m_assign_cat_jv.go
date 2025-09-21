package m_json

import (
	"Spy-Cat-Agency/src/internal/missions/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JsonAssignCatToMissionValidator() gin.HandlerFunc {
	return func(c *gin.Context) {

		var spyCatReq dtos.AssignCatRequest

		if err := c.ShouldBindJSON(&spyCatReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid spy cat data"})
			c.Abort()
			return
		}

		c.Set("spyCatReq", spyCatReq)
		c.Next()
	}
}
