package services

import (
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"Spy-Cat-Agency/src/internal/spycats/dtos"
	"context"
	"net/http"
)

func (s *spyCatServiceImpl) DeleteManySpyCats(ctx context.Context, ids dtos.DeletedIds) error {

	if err := s.deleter.DeleteMany(ctx, ids.Ids); err != nil {
		return error_handler.NewCustomError(http.StatusInternalServerError, "Error deleting many spy cats from the database", err)
	}

	return nil

}
