package interfaces

import (
	"Spy-Cat-Agency/src/internal/spycats/interfaces/handlers"

	"github.com/gin-gonic/gin"
)

func SetUpSpyCatRouter(rg *gin.RouterGroup, h *handlers.SpyCatHandler) {
	spycats := rg.Group("spycats")
	{
		spycats.GET("/", h.GetAllSpyCatsHandler)
		spycats.GET("/:spycatId", h.GetSingleSpyCatHandler)

		spycats.POST("/", h.CreateSpyCatHandler)

		spycats.PATCH("/:spycatId", h.UpdateSpyCatSalaryHandler)

		spycats.DELETE("/", h.DeleteManySpyCatsHandler)
		spycats.DELETE("/:spycatId", h.DeleteSingleSpyCatHandler)
	}
}
