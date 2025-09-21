package handlers

import (
	"Spy-Cat-Agency/src/internal/missions/dtos"
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-faster/errors"
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

	missionId := c.MustGet("missionId").(uuid.UUID)

	spyCatReq := c.MustGet("spyCatReq").(dtos.AssignCatRequest)

	updatedMission, err := h.Service.AssignCatToMission(c.Request.Context(), spyCatReq, missionId)
	if err != nil {
		var customErr *error_handler.CustomError
		if errors.As(err, &customErr) {
			c.JSON(customErr.Code, gin.H{"error": customErr.Message})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, updatedMission)

}
