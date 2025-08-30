package handlers

import (
	"Spy-Cat-Agency/src/internal/spycats/application/services"
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

	var spycatReq dtos.SpyCatCreateRequest

	if err := c.ShouldBindJSON(&spycatReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid spycat data"})
		return
	}

	newSpycat, err := h.Service.CreateSpyCat(c.Request.Context(), spycatReq)
	if err != nil {
		if errors.Is(err, services.ErrorInvalidBreed) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newSpycat)

}
