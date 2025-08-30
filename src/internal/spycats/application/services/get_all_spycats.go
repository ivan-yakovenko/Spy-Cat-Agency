package services

import (
	"Spy-Cat-Agency/src/internal/spycats/dtos"
	"Spy-Cat-Agency/src/internal/spycats/mappers"
	"context"
)

func (s *spyCatServiceImpl) GetAllSpyCats(ctx context.Context) ([]dtos.SpyCatAllResponseDto, error) {

	spycats, err := s.reader.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	return mappers.SpyCatsToDto(spycats), nil

}
