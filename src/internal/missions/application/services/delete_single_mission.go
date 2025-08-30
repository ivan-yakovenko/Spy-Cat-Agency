package services

import (
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"context"
	"log"

	"github.com/go-faster/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

var ErrorMissionAssigned = errors.New("The mission is already assigned to the spy cat")

func (s *missionServiceImpl) DeleteSingleMission(ctx context.Context, id uuid.UUID) error {

	tx, err := s.pool.BeginTx(ctx, pgx.TxOptions{})

	if err != nil {
		return error_handler.ErrorHandler(err, err.Error())
	}

	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			log.Printf("Error rolling back transaction: %v", err)
		}
	}()

	mission, err := s.reader.FindMissionById(ctx, id)

	if err != nil {
		return err
	}

	if mission.SpyCatId != nil && *mission.SpyCatId != uuid.Nil {
		return ErrorMissionAssigned
	}

	if err := s.deleter.DeleteMissionById(ctx, tx, id); err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil

}
