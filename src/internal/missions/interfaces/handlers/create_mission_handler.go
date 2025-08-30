package handlers

import (
	"Spy-Cat-Agency/src/internal/missions/dtos"
	"Spy-Cat-Agency/src/internal/spycats/application/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-faster/errors"
)

// CreateMissionHandler godoc
// @Summary      Create a new mission
// @Description  Creates a new mission with the provided data.
// @Tags         missions
// @Accept       json
// @Produce      json
// @Param        mission body dtos.MissionCreateRequest true "Mission creation data"
// @Success      201 {object} dtos.MissionSingleCreateResponseDto
// @Failure      400 {object} map[string]interface{} "Invalid mission data or targets number"
// @Failure      500 {object} map[string]interface{} "Internal server error"
// @Router       /missions [post]
func (h *MissionHandler) CreateMissionHandler(c *gin.Context) {

	var missionReq dtos.MissionCreateRequest

	if err := c.ShouldBindJSON(&missionReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mission data"})
		return
	}

	newMission, err := h.Service.CreateMission(c.Request.Context(), missionReq)
	if err != nil {
		if errors.Is(err, services.ErrorInvalidBreed) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newMission)

}
