package services

import (
	"Spy-Cat-Agency/src/internal/missions/domain/models"
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"context"
	"log"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

var ErrorTargetCompleted = errors.New("Target is already completed")

func (s *missionServiceImpl) DeleteSingleTarget(ctx context.Context, id uuid.UUID) error {

	tx, err := s.pool.BeginTx(ctx, pgx.TxOptions{})

	if err != nil {
		return error_handler.NewCustomError(http.StatusInternalServerError, "Error starting transaction to delete a target", err)
	}

	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			log.Printf("Error rolling back deleting a target transaction: %v", err)
		}
	}()

	target, err := s.reader.FindTargetById(ctx, id)

	if err != nil {
		return error_handler.NewCustomError(http.StatusInternalServerError, "Error getting single target data from the database", err)
	}

	if target.CompleteState == models.Completed {
		return error_handler.NewCustomError(http.StatusBadRequest, "Can not delete a target, target is already completed", ErrorTargetCompleted)
	}

	if err := s.deleter.DeleteTargetById(ctx, tx, id); err != nil {
		return error_handler.NewCustomError(http.StatusInternalServerError, "Error deleting a target from a database", ErrorMissionAssigned)
	}

	if err := tx.Commit(ctx); err != nil {
		return error_handler.NewCustomError(http.StatusInternalServerError, "Error commiting transaction to delete a single target", err)
	}

	return nil

}
