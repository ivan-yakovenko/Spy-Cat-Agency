package services

import (
	"Spy-Cat-Agency/src/internal/missions/domain/models"
	"Spy-Cat-Agency/src/internal/missions/dtos"
	"Spy-Cat-Agency/src/internal/missions/mappers"
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"context"
	"log"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

var ErrorNotesCantBeModified = errors.New("Notes can't be modified")

func (s *missionServiceImpl) UpdateTarget(ctx context.Context, targetReq dtos.TargetUpdateRequest, missionId uuid.UUID, targetId uuid.UUID) (dtos.MissionSingleResponseDto, error) {

	tx, err := s.pool.BeginTx(ctx, pgx.TxOptions{})

	if err != nil {
		return dtos.MissionSingleResponseDto{}, error_handler.NewCustomError(http.StatusInternalServerError, "Error starting transaction to update single target", err)
	}

	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			log.Printf("Error rolling back updating single target transaction: %v", err)
		}
	}()

	target, err := s.reader.FindTargetById(ctx, targetId)

	if err != nil {
		return dtos.MissionSingleResponseDto{}, error_handler.NewCustomError(http.StatusInternalServerError, "Error getting single target data from the database", err)
	}

	if target.CompleteState == models.Completed {
		targetReq.Notes = ""
		return dtos.MissionSingleResponseDto{}, error_handler.NewCustomError(http.StatusBadRequest, "Notes can not be modified, completion state of the target is already marked 'Completed'", ErrorNotesCantBeModified)
	}

	updatedTarget := mappers.TargetUpdateWithDto(targetReq, target)

	if _, err := s.updater.UpdateTarget(ctx, tx, updatedTarget); err != nil {
		return dtos.MissionSingleResponseDto{}, error_handler.NewCustomError(http.StatusInternalServerError, "Error updating single target in database", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return dtos.MissionSingleResponseDto{}, error_handler.NewCustomError(http.StatusInternalServerError, "Error commiting transaction to update single target", err)
	}

	mission, updTargets, spyCat, err := s.reader.FindMissionById(ctx, missionId)

	if err != nil {
		return dtos.MissionSingleResponseDto{}, error_handler.NewCustomError(http.StatusInternalServerError, "Error getting single mission data from the database", err)
	}

	return mappers.MissionSingleToDto(mission, spyCat, updTargets), nil

}
