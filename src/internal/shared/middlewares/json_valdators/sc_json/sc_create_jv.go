package sc_json

import (
	"Spy-Cat-Agency/src/internal/spycats/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JsonCreateSpyCatValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var spycatReq dtos.SpyCatCreateRequest

		if err := c.ShouldBindJSON(&spycatReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid spycat data"})
			c.Abort()
			return
		}
		c.Set("spycatReq", spycatReq)
		c.Next()
	}
}
