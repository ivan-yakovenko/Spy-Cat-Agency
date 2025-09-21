package sc_json

import (
	"Spy-Cat-Agency/src/internal/spycats/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JsonDeleteManySpyCatValidator() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ids dtos.DeletedIds

		if err := c.ShouldBindJSON(&ids); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid spycat ids"})
			c.Abort()
			return
		}

		c.Set("ids", ids)
		c.Next()
	}
}
