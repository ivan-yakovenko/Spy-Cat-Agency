package handlers

import (
	"Spy-Cat-Agency/src/internal/missions/application/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-faster/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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

	targetIdStr := c.Param("targetId")

	targetId, err := uuid.Parse(targetIdStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid target Id"})
		return
	}

	if err := h.Service.DeleteSingleTarget(c.Request.Context(), targetId); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Mission not found"})
			return
		}
		if errors.Is(err, services.ErrorTargetCompleted) {
			c.JSON(http.StatusBadRequest, gin.H{"error": services.ErrorTargetCompleted.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})

}
