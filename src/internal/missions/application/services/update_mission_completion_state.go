package services

import (
	"Spy-Cat-Agency/src/internal/missions/domain/models"
	"Spy-Cat-Agency/src/internal/missions/dtos"
	"Spy-Cat-Agency/src/internal/missions/mappers"
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	models2 "Spy-Cat-Agency/src/internal/spycats/domain/models"
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (s *missionServiceImpl) UpdateMissionCompletionState(ctx context.Context, completeReq dtos.CompletionStateRequest, id uuid.UUID) (dtos.MissionSingleResponseDto, error) {

	tx, err := s.pool.BeginTx(ctx, pgx.TxOptions{})

	if err != nil {
		return dtos.MissionSingleResponseDto{}, error_handler.ErrorHandler(err, err.Error())
	}

	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			log.Printf("Error rolling back transaction: %v", err)
		}
	}()

	mission := &models.Mission{
		Id:            id,
		CompleteState: completeReq.CompleteState,
	}

	updatedMission, err := s.updater.UpdateMission(ctx, tx, mission)
	if err != nil {
		return dtos.MissionSingleResponseDto{}, err
	}

	var spyCat *models2.SpyCat

	if mission.SpyCatId != nil {
		spyCat, err = s.spyCatReader.FindById(ctx, *mission.SpyCatId)
		if err != nil {
			return dtos.MissionSingleResponseDto{}, err
		}

	} else {
		spyCat = &models2.SpyCat{}
	}

	updTargets, err := s.reader.FindMissionTargetsById(ctx, mission.Id)
	if err != nil {
		return dtos.MissionSingleResponseDto{}, err
	}

	if err := tx.Commit(ctx); err != nil {
		return dtos.MissionSingleResponseDto{}, err
	}

	return mappers.MissionSingleToDto(updatedMission, spyCat, updTargets), nil

}
