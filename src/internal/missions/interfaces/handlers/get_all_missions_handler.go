package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, spycats)

}
