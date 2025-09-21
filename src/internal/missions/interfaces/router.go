package interfaces

import (
	"Spy-Cat-Agency/src/internal/missions/interfaces/handlers"
	"Spy-Cat-Agency/src/internal/shared/middlewares/id_validators"
	"Spy-Cat-Agency/src/internal/shared/middlewares/json_valdators/m_json"
	"Spy-Cat-Agency/src/internal/shared/middlewares/json_valdators/t_json"

	"github.com/gin-gonic/gin"
)

func SetUpMissionRouter(rg *gin.RouterGroup, h *handlers.MissionHandler) {
	missions := rg.Group("missions")
	{
		missions.GET("/", h.GetAllMissionsHandler)
		missions.GET("/:missionId", id_validators.MissionIdValidator(), h.GetSingleMissionHandler)

		missions.POST("/", m_json.JsonCreateMissionValidator(), h.CreateMissionHandler)
		missions.POST("/:missionId/spycats", id_validators.MissionIdValidator(), m_json.JsonAssignCatToMissionValidator(), h.AssignCatToMissionHandler)
		missions.POST("/:missionId/targets", id_validators.MissionIdValidator(), t_json.JsonCreateTargetValidator(), h.CreateTargetHandler)

		missions.PATCH("/:missionId/completion-state", id_validators.MissionIdValidator(), m_json.JsonUpdateMissionValidator(), h.UpdateMissionCompletionStateHandler)
		missions.PATCH("/:missionId/targets/:targetId", id_validators.MissionIdValidator(), id_validators.TargetIdValidator(), t_json.JsonUpdateTargetValidator(), h.UpdateTargetHandler)

		missions.DELETE("/:missionId", id_validators.MissionIdValidator(), h.DeleteSingleMissionHandler)
		missions.DELETE("/:missionId/targets/:targetId", id_validators.MissionIdValidator(), id_validators.TargetIdValidator(), h.DeleteSingleTargetHandler)

	}
}
