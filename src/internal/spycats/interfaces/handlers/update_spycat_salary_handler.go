package handlers

import (
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"Spy-Cat-Agency/src/internal/spycats/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-faster/errors"
	"github.com/google/uuid"
)

// UpdateSpyCatSalaryHandler godoc
// @Summary      Update spy cat salary
// @Description  Updates the salary of a spy cat by its I.
// @Tags         spycats
// @Accept       json
// @Produce      json
// @Param        spycatId path string true "Spycat Id (UUID)"
// @Param        salary body dtos.SalaryRequest true "New salary data"
// @Success      200 {object} dtos.SpyCatSingleResponseDto
// @Failure      400 {object} map[string]interface{} "Invalid spycat Id or salary data"
// @Failure      404 {object} map[string]interface{} "Spy cat not found"
// @Failure      500 {object} map[string]interface{} "Internal server error"
// @Router       /spycats/{spycatId}/salary [patch]
func (h *SpyCatHandler) UpdateSpyCatSalaryHandler(c *gin.Context) {

	id := c.MustGet("spycatId").(uuid.UUID)
	salaryReq := c.MustGet("salaryReq").(dtos.SalaryRequest)

	spycat, err := h.Service.UpdateSpyCatSalary(c.Request.Context(), salaryReq, id)
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
