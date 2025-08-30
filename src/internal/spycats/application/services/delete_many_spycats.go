package services

import (
	"Spy-Cat-Agency/src/internal/spycats/dtos"
	"context"
)

func (s *spyCatServiceImpl) DeleteManySpyCats(ctx context.Context, ids dtos.DeletedIds) error {

	if err := s.deleter.DeleteMany(ctx, ids.Ids); err != nil {
		return err
	}

	return nil

}
