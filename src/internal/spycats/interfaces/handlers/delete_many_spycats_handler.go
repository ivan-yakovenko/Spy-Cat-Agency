package handlers

import (
	"Spy-Cat-Agency/src/internal/spycats/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-faster/errors"
	"github.com/jackc/pgx/v5"
)

// DeleteManySpyCatsHandler godoc
// @Summary      Delete many spy cats
// @Description  Deletes multiple spy cats by their Ids.
// @Tags         spycats
// @Accept       json
// @Produce      json
// @Param        request body dtos.DeletedIds true "Ids to delete"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{} "Invalid spycat ids"
// @Failure      404 {object} map[string]interface{} "Spy cat not found"
// @Failure      500 {object} map[string]interface{} "Internal server error"
// @Router       /spycats [delete]
func (h *SpyCatHandler) DeleteManySpyCatsHandler(c *gin.Context) {

	var ids dtos.DeletedIds

	if err := c.ShouldBindJSON(&ids); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid spycat ids"})
		return
	}

	if err := h.Service.DeleteManySpyCats(c.Request.Context(), ids); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Spy cat not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})

}
