package handlers

import (
	"Spy-Cat-Agency/src/internal/missions/application/services"
	"Spy-Cat-Agency/src/internal/missions/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-faster/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// UpdateTargetHandler godoc
// @Summary      Update mission target
// @Description  Updates a mission target by its Id. Notes cannot be modified if the target is completed.
// @Tags         missions
// @Accept       json
// @Produce      json
// @Param        missionId path string true "Mission Id (UUID)"
// @Param        targetId path string true "Target Id (UUID)"
// @Param        target body dtos.TargetUpdateRequest true "Target update data"
// @Success      200 {object} dtos.MissionSingleResponseDto
// @Failure      400 {object} map[string]interface{} "Invalid mission Id, target Id, or notes can't be modified"
// @Failure      404 {object} map[string]interface{} "Target not found"
// @Failure      500 {object} map[string]interface{} "Internal server error"
// @Router       /missions/{missionId}/targets/{targetId} [patch]
func (h *MissionHandler) UpdateTargetHandler(c *gin.Context) {

	missionIdStr := c.Param("missionId")

	missionId, err := uuid.Parse(missionIdStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mission Id"})
		return
	}

	targetIdStr := c.Param("targetId")

	targetId, err := uuid.Parse(targetIdStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid target Id"})
		return
	}

	var targetReq dtos.TargetUpdateRequest

	if err := c.ShouldBindJSON(&targetReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid completion state data"})
		return
	}

	target, err := h.Service.UpdateTarget(c.Request.Context(), targetReq, missionId, targetId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Target not found"})
			return
		}
		if errors.Is(err, services.ErrorNotesCantBeModified) {
			c.JSON(http.StatusBadRequest, gin.H{"error": services.ErrorNotesCantBeModified.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, target)

}
