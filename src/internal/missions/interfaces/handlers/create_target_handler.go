package handlers

import (
	"Spy-Cat-Agency/src/internal/missions/dtos"
	"Spy-Cat-Agency/src/internal/spycats/application/services"
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

	idStr := c.Param("missionId")

	id, err := uuid.Parse(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mission Id"})
		return
	}

	var targetReq dtos.MissionTargetRequest

	if err := c.ShouldBindJSON(&targetReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid target data"})
		return
	}

	newTarget, err := h.Service.CreateTarget(c.Request.Context(), targetReq, id)
	if err != nil {
		if errors.Is(err, services.ErrorInvalidBreed) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newTarget)

}
