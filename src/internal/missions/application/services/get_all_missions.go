package services

import (
	"Spy-Cat-Agency/src/internal/missions/dtos"
	"Spy-Cat-Agency/src/internal/missions/mappers"
	"context"

	"github.com/google/uuid"
)

func (s *missionServiceImpl) GetAllMissions(ctx context.Context) ([]dtos.MissionAllResponseDto, error) {

	missions, err := s.reader.FindAllMissions(ctx)

	if err != nil {
		return nil, err
	}

	spyCats := make(map[uuid.UUID]string)
	targetsNames := make(map[uuid.UUID][]string)

	for _, mission := range missions {

		if mission.SpyCatId != nil {
			spyCat, err := s.spyCatReader.FindById(ctx, *mission.SpyCatId)
			if err != nil {
				return nil, err
			}

			spyCats[*mission.SpyCatId] = spyCat.Name
		} else {
			spyCats[uuid.Nil] = ""
		}

		targets, err := s.reader.FindMissionTargetsById(ctx, mission.Id)
		if err != nil {
			return nil, err
		}

		targetsNames[mission.Id] = mappers.TargetsNamesToDto(targets)

	}

	return mappers.MissionsToDto(missions, spyCats, targetsNames), nil

}
