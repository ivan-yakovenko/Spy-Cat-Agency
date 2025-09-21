package domain

import (
	"Spy-Cat-Agency/src/internal/missions/domain/models"
	models2 "Spy-Cat-Agency/src/internal/spycats/domain/models"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type MissionReader interface {
	FindAllMissions(ctx context.Context) (map[uuid.UUID]*models.MissionDetails, error)
	FindMissionById(ctx context.Context, id uuid.UUID) (*models.Mission, []models.Target, *models2.SpyCat, error)
	FindTargetById(ctx context.Context, targetId uuid.UUID) (*models.Target, error)
	FindMissionTargetsById(ctx context.Context, missionId uuid.UUID) ([]models.Target, error)
}

type MissionWriter interface {
	Create(ctx context.Context, tx pgx.Tx, newMission *models.Mission, newTargets []models.Target) (*models.Mission, error)
	AssignCatToMission(ctx context.Context, tx pgx.Tx, mission *models.Mission) (*models.Mission, error)
	CreateTarget(ctx context.Context, tx pgx.Tx, target *models.Target) (*models.Target, error)
}

type MissionUpdater interface {
	UpdateMission(ctx context.Context, tx pgx.Tx, updatedMission *models.Mission) (*models.Mission, error)
	UpdateTarget(ctx context.Context, tx pgx.Tx, updatedTarget *models.Target) (*models.Target, error)
}

type MissionDeleter interface {
	DeleteMissionById(ctx context.Context, tx pgx.Tx, id uuid.UUID) error
	DeleteTargetById(ctx context.Context, tx pgx.Tx, id uuid.UUID) error
}
