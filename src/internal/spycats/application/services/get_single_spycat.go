package services

import (
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"Spy-Cat-Agency/src/internal/spycats/dtos"
	"Spy-Cat-Agency/src/internal/spycats/mappers"
	"context"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/google/uuid"
)

func (s *spyCatServiceImpl) GetSingleSpyCat(ctx context.Context, id uuid.UUID) (dtos.SpyCatSingleResponseDto, error) {

	spycat, err := s.reader.FindById(ctx, id)
	if err != nil {
		var customErr *error_handler.CustomError
		if errors.As(err, &customErr) {
			return dtos.SpyCatSingleResponseDto{}, customErr
		}
		return dtos.SpyCatSingleResponseDto{}, error_handler.NewCustomError(http.StatusInternalServerError, "Error getting single spy cat from the database", err)
	}

	return mappers.SpyCatSingleToDto(spycat), nil

}
