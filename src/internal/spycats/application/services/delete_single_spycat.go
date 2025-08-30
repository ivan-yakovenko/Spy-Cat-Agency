package services

import (
	"context"

	"github.com/google/uuid"
)

func (s *spyCatServiceImpl) DeleteSingleSpyCat(ctx context.Context, id uuid.UUID) error {

	if err := s.deleter.DeleteById(ctx, id); err != nil {
		return err
	}

	return nil

}
