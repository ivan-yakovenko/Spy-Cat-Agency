package services

import (
	"Spy-Cat-Agency/src/internal/missions/dtos"
	"Spy-Cat-Agency/src/internal/missions/mappers"
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (s *missionServiceImpl) AssignCatToMission(ctx context.Context, spyCatReq dtos.AssignCatRequest, id uuid.UUID) (dtos.MissionSingleResponseDto, error) {

	tx, err := s.pool.BeginTx(ctx, pgx.TxOptions{})

	if err != nil {
		return dtos.MissionSingleResponseDto{}, error_handler.NewCustomError(http.StatusInternalServerError, "Error starting transaction to assign cat to a mission", err)
	}

	defer func() {
		if err := tx.Rollback(ctx); err != nil {
			log.Printf("Error rolling back assigning cat to a mission transaction: %v", err)
		}
	}()

	mission, _, _, err := s.reader.FindMissionById(ctx, id)

	if err != nil {
		return dtos.MissionSingleResponseDto{}, error_handler.NewCustomError(http.StatusInternalServerError, "Error getting single mission related data from the database", err)
	}

	mission.SpyCatId = &spyCatReq.SpyCatId
	mission.UpdatedAt = time.Now()

	if _, err = s.writer.AssignCatToMission(ctx, tx, mission); err != nil {
		return dtos.MissionSingleResponseDto{}, error_handler.NewCustomError(http.StatusInternalServerError, "Error updating mission with new spy cat in database", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return dtos.MissionSingleResponseDto{}, error_handler.NewCustomError(http.StatusInternalServerError, "Error commiting transaction to assign cat to a mission", err)
	}

	updMission, updTargets, newSpyCat, err := s.reader.FindMissionById(ctx, id)

	if err != nil {
		return dtos.MissionSingleResponseDto{}, error_handler.NewCustomError(http.StatusInternalServerError, err.Error(), err)
	}

	return mappers.MissionSingleToDto(updMission, newSpyCat, updTargets), nil

}
