package handlers

import (
	"Spy-Cat-Agency/src/internal/missions/dtos"
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-faster/errors"
	"github.com/google/uuid"
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
// @Router       /missions/{missionId}/completion-state [patch]
func (h *MissionHandler) UpdateMissionCompletionStateHandler(c *gin.Context) {

	missionId := c.MustGet("missionId").(uuid.UUID)

	completeReq := c.MustGet("completeReq").(dtos.CompletionStateRequest)

	mission, err := h.Service.UpdateMissionCompletionState(c.Request.Context(), completeReq, missionId)
	if err != nil {
		var customErr *error_handler.CustomError
		if errors.As(err, &customErr) {
			c.JSON(customErr.Code, gin.H{"error": customErr.Message})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, mission)

}
