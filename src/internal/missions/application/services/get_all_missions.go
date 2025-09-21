package services

import (
	"Spy-Cat-Agency/src/internal/missions/dtos"
	"Spy-Cat-Agency/src/internal/missions/mappers"
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"context"
	"net/http"
)

func (s *missionServiceImpl) GetAllMissions(ctx context.Context) ([]dtos.MissionAllResponseDto, error) {

	missions, err := s.reader.FindAllMissions(ctx)

	if err != nil {
		return nil, error_handler.NewCustomError(http.StatusInternalServerError, "Error getting all missions data from the database", err)
	}

	return mappers.MissionsToDto(missions), nil

}
