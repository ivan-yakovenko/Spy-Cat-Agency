package handlers

import (
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-faster/errors"
)

// GetAllMissionsHandler godoc
// @Summary      Get all missions
// @Description  Retrieves a list of all missions with their spy cats and targets.
// @Tags         missions
// @Accept       json
// @Produce      json
// @Success      200 {array} dtos.MissionAllResponseDto
// @Failure      500 {object} map[string]interface{} "Internal server error"
// @Router       /missions [get]
func (h *MissionHandler) GetAllMissionsHandler(c *gin.Context) {

	spycats, err := h.Service.GetAllMissions(c.Request.Context())

	if err != nil {
		var customErr *error_handler.CustomError
		if errors.As(err, &customErr) {
			c.JSON(customErr.Code, gin.H{"error": customErr.Message})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, spycats)

}
