package services

import (
	"Spy-Cat-Agency/src/internal/missions/domain/models"
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"context"
	"log"

	"github.com/go-faster/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

var ErrorTargetCompleted = errors.New("Target is already completed")

func (s *missionServiceImpl) DeleteSingleTarget(ctx context.Context, id uuid.UUID) error {

	tx, err := s.pool.BeginTx(ctx, pgx.TxOptions{})

	if err != nil {
		return error_handler.ErrorHandler(err, err.Error())
	}

	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			log.Printf("Error rolling back transaction: %v", err)
		}
	}()

	target, err := s.reader.FindTargetById(ctx, id)

	if err != nil {
		return err
	}

	if target.CompleteState == models.Completed {
		return ErrorTargetCompleted
	}

	if err := s.deleter.DeleteTargetById(ctx, tx, id); err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil

}
