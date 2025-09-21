package services

import (
	"Spy-Cat-Agency/src/internal/missions/domain/models"
	"Spy-Cat-Agency/src/internal/missions/dtos"
	"Spy-Cat-Agency/src/internal/missions/mappers"
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"context"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (s *missionServiceImpl) UpdateMissionCompletionState(ctx context.Context, completeReq dtos.CompletionStateRequest, id uuid.UUID) (dtos.MissionSingleResponseDto, error) {

	tx, err := s.pool.BeginTx(ctx, pgx.TxOptions{})

	if err != nil {
		return dtos.MissionSingleResponseDto{}, error_handler.NewCustomError(http.StatusInternalServerError, "Error starting transaction to update single mission completion state", err)
	}

	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			log.Printf("Error rolling back updating single mission completion state transaction: %v", err)
		}
	}()

	mission := &models.Mission{
		Id:            id,
		CompleteState: completeReq.CompleteState,
	}

	if _, err = s.updater.UpdateMission(ctx, tx, mission); err != nil {
		return dtos.MissionSingleResponseDto{}, error_handler.NewCustomError(http.StatusInternalServerError, "Error updating single mission completion state in database", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return dtos.MissionSingleResponseDto{}, error_handler.NewCustomError(http.StatusInternalServerError, "Error commiting transaction to update single mission completion state", err)
	}

	updMission, targets, spyCat, err := s.reader.FindMissionById(ctx, id)

	if err != nil {
		return dtos.MissionSingleResponseDto{}, error_handler.NewCustomError(http.StatusInternalServerError, "Error getting single mission data from the database", err)
	}

	return mappers.MissionSingleToDto(updMission, spyCat, targets), nil

}
