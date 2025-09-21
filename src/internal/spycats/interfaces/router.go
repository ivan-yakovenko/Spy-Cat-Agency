package interfaces

import (
	"Spy-Cat-Agency/src/internal/shared/middlewares/id_validators"
	"Spy-Cat-Agency/src/internal/shared/middlewares/json_valdators/sc_json"
	"Spy-Cat-Agency/src/internal/spycats/interfaces/handlers"

	"github.com/gin-gonic/gin"
)

func SetUpSpyCatRouter(rg *gin.RouterGroup, h *handlers.SpyCatHandler) {
	spycats := rg.Group("spycats")
	{
		spycats.GET("/", h.GetAllSpyCatsHandler)
		spycats.GET("/:spycatId", id_validators.SpyCatIdValidator(), h.GetSingleSpyCatHandler)

		spycats.POST("/", sc_json.JsonCreateSpyCatValidator(), h.CreateSpyCatHandler)

		spycats.PATCH("/:spycatId/salary", id_validators.SpyCatIdValidator(), sc_json.JsonUpdateSpyCatValidator(), h.UpdateSpyCatSalaryHandler)

		spycats.DELETE("/", sc_json.JsonDeleteManySpyCatValidator(), h.DeleteManySpyCatsHandler)
		spycats.DELETE("/:spycatId", id_validators.SpyCatIdValidator(), h.DeleteSingleSpyCatHandler)
	}
}
