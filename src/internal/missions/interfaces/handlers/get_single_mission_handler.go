package handlers

import (
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-faster/errors"
	"github.com/google/uuid"
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

	missionId := c.MustGet("missionId").(uuid.UUID)

	mission, err := h.Service.GetSingleMission(c.Request.Context(), missionId)
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
