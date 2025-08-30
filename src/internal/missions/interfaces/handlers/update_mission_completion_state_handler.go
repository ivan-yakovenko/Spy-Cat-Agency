package handlers

import (
	"Spy-Cat-Agency/src/internal/missions/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-faster/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// UpdateMissionCompletionStateHandler godoc
// @Summary      Update mission completion state
// @Description  Updates the completion state of a mission by its Id.
// @Tags         missions
// @Accept       json
// @Produce      json
// @Param        missionId path string true "Mission Id (UUID)"
// @Param        completionState body dtos.CompletionStateRequest true "Completion state data"
// @Success      200 {object} dtos.MissionSingleResponseDto
// @Failure      400 {object} map[string]interface{} "Invalid mission Id or completion state data"
// @Failure      404 {object} map[string]interface{} "Mission not found"
// @Failure      500 {object} map[string]interface{} "Internal server error"
// @Router       /missions/{missionId} [patch]
func (h *MissionHandler) UpdateMissionCompletionStateHandler(c *gin.Context) {

	idStr := c.Param("missionId")

	id, err := uuid.Parse(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mission Id"})
		return
	}

	var completeReq dtos.CompletionStateRequest

	if err := c.ShouldBindJSON(&completeReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid completion state data"})
		return
	}

	mission, err := h.Service.UpdateMissionCompletionState(c.Request.Context(), completeReq, id)
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
