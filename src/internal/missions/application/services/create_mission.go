package services

import (
	"Spy-Cat-Agency/src/internal/missions/dtos"
	"Spy-Cat-Agency/src/internal/missions/mappers"
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"context"
	"log"

	"github.com/go-faster/errors"
	"github.com/jackc/pgx/v5"
)

const MAXTARGETS = 3
const MINTARGET = 1

var ErrorInvalidTargetsNumber = errors.New("Invalid targets number")

func (s *missionServiceImpl) CreateMission(ctx context.Context, missionReq dtos.MissionCreateRequest) (dtos.MissionSingleCreateResponseDto, error) {

	if len(missionReq.Targets) > MAXTARGETS || len(missionReq.Targets) < MINTARGET {
		return dtos.MissionSingleCreateResponseDto{}, ErrorInvalidTargetsNumber
	}

	tx, err := s.pool.BeginTx(ctx, pgx.TxOptions{})

	if err != nil {
		return dtos.MissionSingleCreateResponseDto{}, error_handler.ErrorHandler(err, err.Error())
	}

	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			log.Printf("Error rolling back transaction: %v", err)
		}
	}()

	newMission := mappers.CreateDtoToMission(missionReq)

	newMissionTargets := mappers.CreateTargetDtoToMission(missionReq, newMission.Id)

	createdMission, err := s.writer.Create(ctx, tx, newMission, newMissionTargets)

	if err != nil {
		return dtos.MissionSingleCreateResponseDto{}, err
	}

	if err := tx.Commit(ctx); err != nil {
		return dtos.MissionSingleCreateResponseDto{}, err
	}

	return mappers.MissionCreateSingleToDto(createdMission, newMissionTargets), nil
}
