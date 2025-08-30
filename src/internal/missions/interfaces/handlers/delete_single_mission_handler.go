package handlers

import (
	"Spy-Cat-Agency/src/internal/missions/application/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-faster/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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

	idStr := c.Param("missionId")

	id, err := uuid.Parse(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mission Id"})
		return
	}

	if err := h.Service.DeleteSingleMission(c.Request.Context(), id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Mission not found"})
			return
		}
		if errors.Is(err, services.ErrorMissionAssigned) {
			c.JSON(http.StatusBadRequest, gin.H{"error": services.ErrorMissionAssigned.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})

}
