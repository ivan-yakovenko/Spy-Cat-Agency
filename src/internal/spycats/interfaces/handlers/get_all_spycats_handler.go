package handlers

import (
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-faster/errors"
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
		var customErr *error_handler.CustomError
		if errors.As(err, &customErr) {
			c.JSON(customErr.Code, gin.H{"error": customErr.Message})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, spycats)

}
