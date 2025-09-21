package handlers

import (
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"Spy-Cat-Agency/src/internal/spycats/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-faster/errors"
)

// CreateSpyCatHandler godoc
// @Summary      Create a new spycat
// @Description  Creates a new spycat with the provided data.
// @Tags         spycats
// @Accept       json
// @Produce      json
// @Param        spycat body dtos.SpyCatCreateRequest true "Spycat data"
// @Success      201 {object} dtos.SpyCatSingleResponseDto
// @Failure      400 {object} map[string]interface{} "Invalid spycat data or breed"
// @Failure      500 {object} map[string]interface{} "Internal server error"
// @Router       /spycats [post]
func (h *SpyCatHandler) CreateSpyCatHandler(c *gin.Context) {

	spycatReq := c.MustGet("spycatReq").(dtos.SpyCatCreateRequest)

	newSpycat, err := h.Service.CreateSpyCat(c.Request.Context(), spycatReq)
	if err != nil {
		var customErr *error_handler.CustomError
		if errors.As(err, &customErr) {
			c.JSON(customErr.Code, gin.H{"error": customErr.Message})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, newSpycat)

}
