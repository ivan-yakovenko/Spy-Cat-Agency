package handlers

import (
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-faster/errors"
	"github.com/google/uuid"
)

// DeleteSingleTargetHandler godoc
// @Summary      Delete a mission target
// @Description  Deletes a mission target by its Id. Fails if the target is already completed.
// @Tags         missions
// @Accept       json
// @Produce      json
// @Param        missionId path string true "Mission Id (UUID)"
// @Param        targetId path string true "Target Id (UUID)"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{} "Invalid target Id"
// @Failure      404 {object} map[string]interface{} "Target not found or already completed"
// @Failure      500 {object} map[string]interface{} "Internal server error"
// @Router       /missions/{missionId}/targets/{targetId} [delete]
func (h *MissionHandler) DeleteSingleTargetHandler(c *gin.Context) {

	targetId := c.MustGet("targetId").(uuid.UUID)

	if err := h.Service.DeleteSingleTarget(c.Request.Context(), targetId); err != nil {
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
