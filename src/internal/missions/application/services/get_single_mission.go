package services

import (
	"Spy-Cat-Agency/src/internal/missions/dtos"
	"Spy-Cat-Agency/src/internal/missions/mappers"
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"context"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/google/uuid"
)

func (s *missionServiceImpl) GetSingleMission(ctx context.Context, id uuid.UUID) (dtos.MissionSingleResponseDto, error) {

	mission, targets, spycat, err := s.reader.FindMissionById(ctx, id)

	if err != nil {
		var customErr *error_handler.CustomError
		if errors.As(err, &customErr) {
			return dtos.MissionSingleResponseDto{}, customErr
		}
		return dtos.MissionSingleResponseDto{}, error_handler.NewCustomError(http.StatusInternalServerError, "Error getting single mission data from the database", err)
	}

	return mappers.MissionSingleToDto(mission, spycat, targets), nil

}
