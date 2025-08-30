package interfaces

import (
	"Spy-Cat-Agency/src/internal/missions/interfaces/handlers"

	"github.com/gin-gonic/gin"
)

func SetUpMissionRouter(rg *gin.RouterGroup, h *handlers.MissionHandler) {
	missions := rg.Group("missions")
	{
		missions.GET("/", h.GetAllMissionsHandler)
		missions.GET("/:missionId", h.GetSingleMissionHandler)

		missions.POST("/", h.CreateMissionHandler)
		missions.POST("/:missionId/spycats", h.AssignCatToMissionHandler)
		missions.POST("/:missionId/targets", h.CreateTargetHandler)

		missions.PATCH("/:missionId", h.UpdateMissionCompletionStateHandler)
		missions.PATCH("/:missionId/targets/:targetId", h.UpdateTargetHandler)

		missions.DELETE("/:missionId", h.DeleteSingleMissionHandler)
		missions.DELETE("/:missionId/targets/:targetId", h.DeleteSingleTargetHandler)

	}
}
