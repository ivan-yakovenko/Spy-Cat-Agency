package services

import (
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"Spy-Cat-Agency/src/internal/spycats/dtos"
	"Spy-Cat-Agency/src/internal/spycats/mappers"
	"context"
	"net/http"
)

func (s *spyCatServiceImpl) CreateSpyCat(ctx context.Context, spycatReq dtos.SpyCatCreateRequest) (dtos.SpyCatSingleResponseDto, error) {

	breedNames, err := fetchBreedsNames()

	if err != nil {
		return dtos.SpyCatSingleResponseDto{}, error_handler.NewCustomError(http.StatusInternalServerError, "Error fetching breed names for the 3-rd party API", err)
	}

	if !isValidBreed(spycatReq.Breed, breedNames) {
		return dtos.SpyCatSingleResponseDto{}, error_handler.NewCustomError(http.StatusBadRequest, "Error breed name inputted", ErrorInvalidBreed)
	}

	newSpycat := mappers.CreateDtoToSpyCat(spycatReq)

	newSpyCat, err := s.writer.Create(ctx, newSpycat)

	if err != nil {
		return dtos.SpyCatSingleResponseDto{}, error_handler.NewCustomError(http.StatusInternalServerError, "Error creating new spy cat", err)
	}

	return mappers.SpyCatSingleToDto(newSpyCat), nil
}
