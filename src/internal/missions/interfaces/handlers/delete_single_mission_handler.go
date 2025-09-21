package handlers

import (
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-faster/errors"
	"github.com/google/uuid"
)

// DeleteSingleMissionHandler godoc
// @Summary      Delete a mission
// @Description  Deletes a mission by its Id. Fails if the mission is already assigned to a spy cat.
// @Tags         missions
// @Accept       json
// @Produce      json
// @Param        missionId path string true "Mission Id (UUID)"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{} "Invalid mission Id"
// @Failure      404 {object} map[string]interface{} "Mission not found or already assigned"
// @Failure      500 {object} map[string]interface{} "Internal server error"
// @Router       /missions/{missionId} [delete]
func (h *MissionHandler) DeleteSingleMissionHandler(c *gin.Context) {

	missionId := c.MustGet("missionId").(uuid.UUID)

	if err := h.Service.DeleteSingleMission(c.Request.Context(), missionId); err != nil {
		var customErr *error_handler.CustomError
		if errors.As(err, &customErr) {
			c.JSON(customErr.Code, gin.H{"error": customErr.Message})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{})

}
