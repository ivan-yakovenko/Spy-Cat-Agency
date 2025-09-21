package sc_json

import (
	"Spy-Cat-Agency/src/internal/spycats/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JsonUpdateSpyCatValidator() gin.HandlerFunc {
	return func(c *gin.Context) {

		var salaryReq dtos.SalaryRequest

		if err := c.ShouldBindJSON(&salaryReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid spycat salary data"})
			c.Abort()
			return
		}

		c.Set("salaryReq", salaryReq)
		c.Next()
	}
}
