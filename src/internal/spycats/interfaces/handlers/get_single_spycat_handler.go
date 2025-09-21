package handlers

import (
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-faster/errors"
	"github.com/google/uuid"
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

	id := c.MustGet("spycatId").(uuid.UUID)

	spycat, err := h.Service.GetSingleSpyCat(c.Request.Context(), id)
	if err != nil {
		var customErr *error_handler.CustomError
		if errors.As(err, &customErr) {
			c.JSON(customErr.Code, gin.H{"error": customErr.Message})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, spycat)

}
