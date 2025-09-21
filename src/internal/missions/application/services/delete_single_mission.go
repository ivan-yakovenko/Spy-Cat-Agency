package services

import (
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"context"
	"log"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

var ErrorMissionAssigned = errors.New("The mission is already assigned to the spy cat")

func (s *missionServiceImpl) DeleteSingleMission(ctx context.Context, id uuid.UUID) error {

	tx, err := s.pool.BeginTx(ctx, pgx.TxOptions{})

	if err != nil {
		return error_handler.NewCustomError(http.StatusInternalServerError, "Error starting transaction to delete a mission", err)
	}

	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			log.Printf("Error rolling back deleting a mission transaction: %v", err)
		}
	}()

	mission, _, _, err := s.reader.FindMissionById(ctx, id)

	if err != nil {
		return error_handler.NewCustomError(http.StatusInternalServerError, "Error getting single mission related data from the database", err)
	}

	if mission.SpyCatId != nil && *mission.SpyCatId != uuid.Nil {
		return error_handler.NewCustomError(http.StatusBadRequest, "Can not delete a mission, mission is currently assigned to a spy cat", ErrorMissionAssigned)
	}

	if err := s.deleter.DeleteMissionById(ctx, tx, id); err != nil {
		return error_handler.NewCustomError(http.StatusInternalServerError, "Error deleting a mission from a database", ErrorMissionAssigned)
	}

	if err := tx.Commit(ctx); err != nil {
		return error_handler.NewCustomError(http.StatusInternalServerError, "Error commiting transaction to delete a single mission for a mission", err)
	}

	return nil

}
