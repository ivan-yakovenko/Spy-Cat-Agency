package services

import (
	"Spy-Cat-Agency/src/internal/missions/dtos"
	"Spy-Cat-Agency/src/internal/missions/mappers"
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"Spy-Cat-Agency/src/internal/spycats/domain/models"
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

const MAXADDEDTARGETS = 2

func (s *missionServiceImpl) CreateTarget(ctx context.Context, targetReq dtos.MissionTargetRequest, id uuid.UUID) (dtos.MissionSingleResponseDto, error) {

	tx, err := s.pool.BeginTx(ctx, pgx.TxOptions{})

	if err != nil {
		return dtos.MissionSingleResponseDto{}, error_handler.ErrorHandler(err, err.Error())
	}

	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			log.Printf("Error rolling back transaction: %v", err)
		}
	}()

	targets, err := s.reader.FindMissionTargetsById(ctx, id)
	if err != nil {
		return dtos.MissionSingleResponseDto{}, err
	}

	if len(targets) > MAXADDEDTARGETS {
		return dtos.MissionSingleResponseDto{}, ErrorInvalidTargetsNumber
	}

	newTarget := mappers.TargetDtoToTarget(targetReq, id)

	if _, err := s.writer.CreateTarget(ctx, tx, newTarget); err != nil {
		return dtos.MissionSingleResponseDto{}, err
	}

	mission, err := s.reader.FindMissionById(ctx, id)

	if err != nil {
		return dtos.MissionSingleResponseDto{}, err
	}

	var spyCat *models.SpyCat

	if mission.SpyCatId != nil {
		spyCat, err = s.spyCatReader.FindById(ctx, *mission.SpyCatId)
		if err != nil {
			return dtos.MissionSingleResponseDto{}, err
		}

	} else {
		spyCat = &models.SpyCat{}
	}

	if err := tx.Commit(ctx); err != nil {
		return dtos.MissionSingleResponseDto{}, err
	}

	updTargets, err := s.reader.FindMissionTargetsById(ctx, mission.Id)
	if err != nil {
		return dtos.MissionSingleResponseDto{}, err
	}

	return mappers.MissionSingleToDto(mission, spyCat, updTargets), nil
}
