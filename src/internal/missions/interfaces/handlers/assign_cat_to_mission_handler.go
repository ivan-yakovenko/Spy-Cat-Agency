package handlers

import (
	"Spy-Cat-Agency/src/internal/missions/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// AssignCatToMissionHandler godoc
// @Summary      Assign a spy cat to a mission
// @Description  Assigns a spy cat to a mission by mission Id.
// @Tags         missions
// @Accept       json
// @Produce      json
// @Param        missionId path string true "Mission I (UUID)"
// @Param        spyCat body dtos.AssignCatRequest true "Spy cat assignment data"
// @Success      201 {object} dtos.MissionSingleResponseDto
// @Failure      400 {object} map[string]interface{} "Invalid mission Id or spy cat data"
// @Failure      500 {object} map[string]interface{} "Internal server error"
// @Router       /missions/{missionId}/spycats [post]
func (h *MissionHandler) AssignCatToMissionHandler(c *gin.Context) {

	idStr := c.Param("missionId")

	id, err := uuid.Parse(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mission Id"})
		return
	}

	var spyCatReq dtos.AssignCatRequest

	if err := c.ShouldBindJSON(&spyCatReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid spy cat data"})
		return
	}

	updatedMission, err := h.Service.AssignCatToMission(c.Request.Context(), spyCatReq, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, updatedMission)

}
