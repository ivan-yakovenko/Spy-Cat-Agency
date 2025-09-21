package services

import (
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"context"
	"net/http"

	"github.com/google/uuid"
)

func (s *spyCatServiceImpl) DeleteSingleSpyCat(ctx context.Context, id uuid.UUID) error {

	if err := s.deleter.DeleteById(ctx, id); err != nil {
		return error_handler.NewCustomError(http.StatusInternalServerError, "Error deleting single spy cat from the database", err)
	}

	return nil

}
