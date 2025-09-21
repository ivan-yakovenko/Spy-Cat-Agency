package handlers

import (
	"Spy-Cat-Agency/src/internal/missions/dtos"
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-faster/errors"
	"github.com/google/uuid"
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

	missionId := c.MustGet("missionId").(uuid.UUID)

	targetId := c.MustGet("targetId").(uuid.UUID)

	targetReq := c.MustGet("targetReq").(dtos.TargetUpdateRequest)

	target, err := h.Service.UpdateTarget(c.Request.Context(), targetReq, missionId, targetId)
	if err != nil {
		var customErr *error_handler.CustomError
		if errors.As(err, &customErr) {
			c.JSON(customErr.Code, gin.H{"error": customErr.Message})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, target)

}
