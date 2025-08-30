package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllSpyCatsHandler godoc
// @Summary      Get all spy cats
// @Description  Retrieves a list of all spy cats.
// @Tags         spycats
// @Accept       json
// @Produce      json
// @Success      200 {array} dtos.SpyCatAllResponseDto
// @Failure      500 {object} map[string]interface{} "Internal server error"
// @Router       /spycats [get]
func (h *SpyCatHandler) GetAllSpyCatsHandler(c *gin.Context) {

	spycats, err := h.Service.GetAllSpyCats(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, spycats)

}
