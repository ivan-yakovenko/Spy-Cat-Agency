package services

import (
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"Spy-Cat-Agency/src/internal/spycats/dtos"
	"Spy-Cat-Agency/src/internal/spycats/mappers"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-faster/errors"
)

func fetchBreedsNames() ([]dtos.BreedName, error) {

	response, err := http.Get("https://api.thecatapi.com/v1/breeds")

	if err != nil {
		return nil, error_handler.ErrorHandler(err, err.Error())
	}

	defer func() {
		if err := response.Body.Close(); err != nil {
			log.Printf("Error closing response body: %v", err)
		}
	}()

	var breedNames []dtos.BreedName

	err = json.NewDecoder(response.Body).Decode(&breedNames)

	if err != nil {
		return nil, error_handler.ErrorHandler(err, err.Error())
	}

	return breedNames, nil

}

func isValidBreed(breedName string, breedNames []dtos.BreedName) bool {
	for _, breed := range breedNames {
		if breed.Name == breedName {
			return true
		}
	}

	return false
}

var ErrorInvalidBreed = errors.New("Invalid breed name")

func (s *spyCatServiceImpl) CreateSpyCat(ctx context.Context, spycatReq dtos.SpyCatCreateRequest) (dtos.SpyCatSingleResponseDto, error) {

	breedNames, err := fetchBreedsNames()

	if err != nil {
		return dtos.SpyCatSingleResponseDto{}, err
	}

	if !isValidBreed(spycatReq.Breed, breedNames) {
		return dtos.SpyCatSingleResponseDto{}, ErrorInvalidBreed
	}

	newSpycat := mappers.CreateDtoToSpyCat(spycatReq)

	newSpyCat, err := s.writer.Create(ctx, newSpycat)

	if err != nil {
		return dtos.SpyCatSingleResponseDto{}, err
	}

	return mappers.SpyCatSingleToDto(newSpyCat), nil
}
