package services

import (
	"Spy-Cat-Agency/src/internal/missions/dtos"
	"Spy-Cat-Agency/src/internal/missions/mappers"
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"context"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

const MAXADDEDTARGETS = 2

func (s *missionServiceImpl) CreateTarget(ctx context.Context, targetReq dtos.MissionTargetRequest, id uuid.UUID) (dtos.MissionSingleResponseDto, error) {

	tx, err := s.pool.BeginTx(ctx, pgx.TxOptions{})

	if err != nil {
		return dtos.MissionSingleResponseDto{}, error_handler.NewCustomError(http.StatusInternalServerError, "Error starting transaction to create a target", err)
	}

	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			log.Printf("Error rolling back creating a target transaction: %v", err)
		}
	}()

	targets, err := s.reader.FindMissionTargetsById(ctx, id)
	if err != nil {
		return dtos.MissionSingleResponseDto{}, error_handler.NewCustomError(http.StatusInternalServerError, "Error getting targets by a mission from a database", err)
	}

	if len(targets) > MAXADDEDTARGETS {
		return dtos.MissionSingleResponseDto{}, error_handler.NewCustomError(http.StatusBadRequest, "Number of targets = 3", ErrorInvalidTargetsNumber)
	}

	newTarget := mappers.TargetDtoToTarget(targetReq, id)

	if _, err := s.writer.CreateTarget(ctx, tx, newTarget); err != nil {
		return dtos.MissionSingleResponseDto{}, error_handler.NewCustomError(http.StatusInternalServerError, "Error creating a target in the database", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return dtos.MissionSingleResponseDto{}, error_handler.NewCustomError(http.StatusInternalServerError, "Error commiting transaction to create a new target for a mission", err)
	}

	mission, updTargets, spyCat, err := s.reader.FindMissionById(ctx, id)

	if err != nil {
		return dtos.MissionSingleResponseDto{}, error_handler.NewCustomError(http.StatusInternalServerError, err.Error(), err)
	}

	return mappers.MissionSingleToDto(mission, spyCat, updTargets), nil
}
