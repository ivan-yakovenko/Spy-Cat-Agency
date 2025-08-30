package application

import (
	"Spy-Cat-Agency/src/internal/missions/dtos"
	"context"

	"github.com/google/uuid"
)

type MissionService interface {
	GetSingleMission(ctx context.Context, id uuid.UUID) (dtos.MissionSingleResponseDto, error)
	GetAllMissions(ctx context.Context) ([]dtos.MissionAllResponseDto, error)

	CreateMission(ctx context.Context, missionReq dtos.MissionCreateRequest) (dtos.MissionSingleCreateResponseDto, error)
	AssignCatToMission(ctx context.Context, spyCatReq dtos.AssignCatRequest, id uuid.UUID) (dtos.MissionSingleResponseDto, error)
	CreateTarget(ctx context.Context, targetReq dtos.MissionTargetRequest, id uuid.UUID) (dtos.MissionSingleResponseDto, error)

	UpdateMissionCompletionState(ctx context.Context, salaryReq dtos.CompletionStateRequest, id uuid.UUID) (dtos.MissionSingleResponseDto, error)
	UpdateTarget(ctx context.Context, targetReq dtos.TargetUpdateRequest, missionId uuid.UUID, targetId uuid.UUID) (dtos.MissionSingleResponseDto, error)

	DeleteSingleMission(ctx context.Context, id uuid.UUID) error
	DeleteSingleTarget(ctx context.Context, id uuid.UUID) error
}
