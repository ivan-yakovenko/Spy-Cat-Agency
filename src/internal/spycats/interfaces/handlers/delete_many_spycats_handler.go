package handlers

import (
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"Spy-Cat-Agency/src/internal/spycats/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-faster/errors"
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

	ids := c.MustGet("ids").(dtos.DeletedIds)

	if err := h.Service.DeleteManySpyCats(c.Request.Context(), ids); err != nil {
		var customErr *error_handler.CustomError
		if errors.As(err, &customErr) {
			c.JSON(customErr.Code, gin.H{"error": customErr.Message})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{})

}
