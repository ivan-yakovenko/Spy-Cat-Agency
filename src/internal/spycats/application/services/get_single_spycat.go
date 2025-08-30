package services

import (
	"Spy-Cat-Agency/src/internal/spycats/dtos"
	"Spy-Cat-Agency/src/internal/spycats/mappers"
	"context"

	"github.com/google/uuid"
)

func (s *spyCatServiceImpl) GetSingleSpyCat(ctx context.Context, id uuid.UUID) (dtos.SpyCatSingleResponseDto, error) {

	spycat, err := s.reader.FindById(ctx, id)
	if err != nil {
		return dtos.SpyCatSingleResponseDto{}, err
	}

	return mappers.SpyCatSingleToDto(spycat), nil

}
