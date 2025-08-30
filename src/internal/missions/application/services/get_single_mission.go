package services

import (
	"Spy-Cat-Agency/src/internal/missions/dtos"
	"Spy-Cat-Agency/src/internal/missions/mappers"
	"Spy-Cat-Agency/src/internal/spycats/domain/models"
	"context"

	"github.com/google/uuid"
)

func (s *missionServiceImpl) GetSingleMission(ctx context.Context, id uuid.UUID) (dtos.MissionSingleResponseDto, error) {

	mission, err := s.reader.FindMissionById(ctx, id)

	if err != nil {
		return dtos.MissionSingleResponseDto{}, err
	}

	var spyCat *models.SpyCat

	if mission.SpyCatId != nil {
		spyCat, err = s.spyCatReader.FindById(ctx, *mission.SpyCatId)
		if err != nil {
			return dtos.MissionSingleResponseDto{}, err
		}

	} else {
		spyCat = &models.SpyCat{}
	}

	targets, err := s.reader.FindMissionTargetsById(ctx, mission.Id)
	if err != nil {
		return dtos.MissionSingleResponseDto{}, err
	}

	return mappers.MissionSingleToDto(mission, spyCat, targets), nil

}
