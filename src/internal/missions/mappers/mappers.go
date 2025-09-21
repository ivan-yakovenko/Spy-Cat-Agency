package mappers

import (
	"Spy-Cat-Agency/src/internal/missions/domain/models"
	"Spy-Cat-Agency/src/internal/missions/dtos"
	models2 "Spy-Cat-Agency/src/internal/spycats/domain/models"
	"Spy-Cat-Agency/src/internal/spycats/mappers"

	"github.com/google/uuid"
)

func MissionTargetsToDto(targets []models.Target) []dtos.TargetDto {
	targetDtos := make([]dtos.TargetDto, 0, len(targets))

	for _, target := range targets {
		targetDtos = append(targetDtos, dtos.TargetDto{
			Id:            target.Id,
			Name:          target.Name,
			Country:       target.Country,
			Notes:         target.Notes,
			CompleteState: string(target.CompleteState),
			UpdatedAt:     target.UpdatedAt,
		})
	}

	return targetDtos
}

func MissionSingleToDto(mission *models.Mission, spyCat *models2.SpyCat, targets []models.Target) dtos.MissionSingleResponseDto {
	return dtos.MissionSingleResponseDto{
		Id:            mission.Id,
		CompleteState: string(mission.CompleteState),
		SpyCat:        mappers.SpyCatSingleToDto(spyCat),
		Targets:       MissionTargetsToDto(targets),
		UpdatedAt:     mission.UpdatedAt,
	}
}

func MissionsToDto(missions map[uuid.UUID]*models.MissionDetails) []dtos.MissionAllResponseDto {

	missionsResponse := make([]dtos.MissionAllResponseDto, 0, len(missions))

	for _, mission := range missions {

		var spyCatName string

		if mission.SpyCatName != nil {
			spyCatName = *mission.SpyCatName
		} else {
			spyCatName = ""
		}

		missionsResponse = append(missionsResponse, dtos.MissionAllResponseDto{
			Id:            mission.Id,
			CompleteState: string(mission.CompleteState),
			SpyCatName:    spyCatName,
			Targets:       mission.TargetNames,
			UpdatedAt:     mission.UpdatedAt,
		})
	}

	return missionsResponse

}

func CreateDtoToMission(missionReq dtos.MissionCreateRequest) *models.Mission {

	mission := &models.Mission{
		Id:            uuid.New(),
		CompleteState: models.CompleteState(missionReq.CompleteState),
	}

	return mission

}

func CreateTargetDtoToMission(missionReq dtos.MissionCreateRequest, missionId uuid.UUID) []models.Target {

	targets := make([]models.Target, 0)

	for _, target := range missionReq.Targets {
		targets = append(targets, models.Target{
			Id:            uuid.New(),
			MissionId:     missionId,
			Name:          target.Name,
			Country:       target.Country,
			Notes:         target.Notes,
			CompleteState: models.CompleteState(target.CompleteState),
		})
	}

	return targets

}

func MissionCreateSingleToDto(mission *models.Mission, targets []models.Target) dtos.MissionSingleCreateResponseDto {
	return dtos.MissionSingleCreateResponseDto{
		Id:            mission.Id,
		CompleteState: string(mission.CompleteState),
		Targets:       MissionTargetsToDto(targets),
		UpdatedAt:     mission.UpdatedAt,
	}
}

func TargetDtoToTarget(targetReq dtos.MissionTargetRequest, id uuid.UUID) *models.Target {
	return &models.Target{
		Id:            uuid.New(),
		MissionId:     id,
		Name:          targetReq.Name,
		Country:       targetReq.Country,
		Notes:         targetReq.Notes,
		CompleteState: models.CompleteState(targetReq.CompleteState),
	}
}

func TargetUpdateWithDto(targetReq dtos.TargetUpdateRequest, target *models.Target) *models.Target {

	if targetReq.CompleteState != "" {
		target.CompleteState = models.CompleteState(targetReq.CompleteState)
	}

	if targetReq.Notes != "" {
		target.Notes = targetReq.Notes
	}

	return target

}
