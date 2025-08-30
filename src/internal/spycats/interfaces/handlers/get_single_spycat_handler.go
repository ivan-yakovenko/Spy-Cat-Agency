package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-faster/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// GetSingleSpyCatHandler godoc
// @Summary      Get a single spy cat
// @Description  Retrieves a spy cat by its Id.
// @Tags         spycats
// @Accept       json
// @Produce      json
// @Param        spycatId path string true "Spycat Id (UUID)"
// @Success      200 {object} dtos.SpyCatSingleResponseDto
// @Failure      400 {object} map[string]interface{} "Invalid spycat Id"
// @Failure      404 {object} map[string]interface{} "Spy cat not found"
// @Failure      500 {object} map[string]interface{} "Internal server error"
// @Router       /spycats/{spycatId} [get]
func (h *SpyCatHandler) GetSingleSpyCatHandler(c *gin.Context) {

	idStr := c.Param("spycatId")

	id, err := uuid.Parse(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid spycat Id"})
		return
	}

	spycat, err := h.Service.GetSingleSpyCat(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Spy cat not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, spycat)

}
