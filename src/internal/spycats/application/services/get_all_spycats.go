package services

import (
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"Spy-Cat-Agency/src/internal/spycats/dtos"
	"Spy-Cat-Agency/src/internal/spycats/mappers"
	"context"
	"net/http"
)

func (s *spyCatServiceImpl) GetAllSpyCats(ctx context.Context) ([]dtos.SpyCatAllResponseDto, error) {

	spycats, err := s.reader.FindAll(ctx)

	if err != nil {
		return nil, error_handler.NewCustomError(http.StatusInternalServerError, "Error getting all spy cats from the database", err)
	}

	return mappers.SpyCatsToDto(spycats), nil

}
