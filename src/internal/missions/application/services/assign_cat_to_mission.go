package services

import (
	"Spy-Cat-Agency/src/internal/missions/dtos"
	"Spy-Cat-Agency/src/internal/missions/mappers"
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"Spy-Cat-Agency/src/internal/spycats/domain/models"
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (s *missionServiceImpl) AssignCatToMission(ctx context.Context, spyCatReq dtos.AssignCatRequest, id uuid.UUID) (dtos.MissionSingleResponseDto, error) {

	tx, err := s.pool.BeginTx(ctx, pgx.TxOptions{})

	if err != nil {
		return dtos.MissionSingleResponseDto{}, error_handler.ErrorHandler(err, err.Error())
	}

	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			log.Printf("Error rolling back transaction: %v", err)
		}
	}()

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

	mission.SpyCatId = &spyCatReq.SpyCatId
	mission.UpdatedAt = time.Now()

	updatedMission, err := s.writer.AssignCatToMission(ctx, tx, mission)

	if err != nil {
		return dtos.MissionSingleResponseDto{}, err
	}

	if err := tx.Commit(ctx); err != nil {
		return dtos.MissionSingleResponseDto{}, err
	}

	targets, err := s.reader.FindMissionTargetsById(ctx, mission.Id)
	if err != nil {
		return dtos.MissionSingleResponseDto{}, err
	}

	spyCat, err = s.spyCatReader.FindById(ctx, *mission.SpyCatId)
	if err != nil {
		return dtos.MissionSingleResponseDto{}, err
	}

	return mappers.MissionSingleToDto(updatedMission, spyCat, targets), nil

}
