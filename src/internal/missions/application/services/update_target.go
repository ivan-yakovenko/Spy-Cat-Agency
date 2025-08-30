package services

import (
	"Spy-Cat-Agency/src/internal/missions/domain/models"
	"Spy-Cat-Agency/src/internal/missions/dtos"
	"Spy-Cat-Agency/src/internal/missions/mappers"
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	models2 "Spy-Cat-Agency/src/internal/spycats/domain/models"
	"context"
	"log"

	"github.com/go-faster/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

var ErrorNotesCantBeModified = errors.New("Notes can't be modified")

func (s *missionServiceImpl) UpdateTarget(ctx context.Context, targetReq dtos.TargetUpdateRequest, missionId uuid.UUID, targetId uuid.UUID) (dtos.MissionSingleResponseDto, error) {

	tx, err := s.pool.BeginTx(ctx, pgx.TxOptions{})

	if err != nil {
		return dtos.MissionSingleResponseDto{}, error_handler.ErrorHandler(err, err.Error())
	}

	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			log.Printf("Error rolling back transaction: %v", err)
		}
	}()

	target, err := s.reader.FindTargetById(ctx, targetId)

	if err != nil {
		return dtos.MissionSingleResponseDto{}, err
	}

	if target.CompleteState == models.Completed {
		targetReq.Notes = ""
		return dtos.MissionSingleResponseDto{}, ErrorNotesCantBeModified
	}

	updatedTarget := mappers.TargetUpdateWithDto(targetReq, target)

	if _, err := s.updater.UpdateTarget(ctx, tx, updatedTarget); err != nil {
		return dtos.MissionSingleResponseDto{}, err
	}

	mission, err := s.reader.FindMissionById(ctx, missionId)

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

	if err := tx.Commit(ctx); err != nil {
		return dtos.MissionSingleResponseDto{}, err
	}

	updTargets, err := s.reader.FindMissionTargetsById(ctx, mission.Id)
	if err != nil {
		return dtos.MissionSingleResponseDto{}, err
	}

	return mappers.MissionSingleToDto(mission, spyCat, updTargets), nil

}
