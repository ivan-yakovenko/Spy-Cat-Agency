package handlers

import (
	"Spy-Cat-Agency/src/internal/missions/dtos"
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-faster/errors"
	"github.com/google/uuid"
)

// CreateTargetHandler godoc
// @Summary      Add a target to a mission
// @Description  Adds a new target to the specified mission by mission Id.
// @Tags         missions
// @Accept       json
// @Produce      json
// @Param        missionId path string true "Mission Id (UUID)"
// @Param        target body dtos.MissionTargetRequest true "Target data"
// @Success      201 {object} dtos.MissionSingleResponseDto
// @Failure      400 {object} map[string]interface{} "Invalid mission Id or target data"
// @Failure      500 {object} map[string]interface{} "Internal server error"
// @Router       /missions/{missionId}/targets [post]
func (h *MissionHandler) CreateTargetHandler(c *gin.Context) {

	missionId := c.MustGet("missionId").(uuid.UUID)

	targetReq := c.MustGet("targetReq").(dtos.MissionTargetRequest)

	newTarget, err := h.Service.CreateTarget(c.Request.Context(), targetReq, missionId)
	if err != nil {
		var customErr *error_handler.CustomError
		if errors.As(err, &customErr) {
			c.JSON(customErr.Code, gin.H{"error": customErr.Message})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, newTarget)

}
