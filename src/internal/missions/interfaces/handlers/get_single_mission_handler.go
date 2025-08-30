package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-faster/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// GetSingleMissionHandler godoc
// @Summary      Get a single mission
// @Description  Retrieves a mission by its Id, including its spy cat and targets.
// @Tags         missions
// @Accept       json
// @Produce      json
// @Param        missionId path string true "Mission Id (UUID)"
// @Success      200 {object} dtos.MissionSingleResponseDto
// @Failure      400 {object} map[string]interface{} "Invalid mission Id"
// @Failure      404 {object} map[string]interface{} "Mission not found"
// @Failure      500 {object} map[string]interface{} "Internal server error"
// @Router       /missions/{missionId} [get]
func (h *MissionHandler) GetSingleMissionHandler(c *gin.Context) {

	idStr := c.Param("missionId")

	id, err := uuid.Parse(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mission Id"})
		return
	}

	mission, err := h.Service.GetSingleMission(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Mission not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mission)

}
