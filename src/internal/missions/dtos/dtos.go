package dtos

import (
	"Spy-Cat-Agency/src/internal/missions/domain/models"
	"Spy-Cat-Agency/src/internal/spycats/dtos"
	"time"

	"github.com/google/uuid"
)

type TargetDto struct {
	Id            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Country       string    `json:"country"`
	Notes         string    `json:"notes"`
	CompleteState string    `json:"complete_state"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type MissionSingleResponseDto struct {
	Id            uuid.UUID                    `json:"id"`
	CompleteState string                       `json:"complete_state"`
	SpyCat        dtos.SpyCatSingleResponseDto `json:"spycat"`
	Targets       []TargetDto                  `json:"targets"`
	UpdatedAt     time.Time                    `json:"updated_at"`
}

type MissionAllResponseDto struct {
	Id            uuid.UUID `json:"id"`
	CompleteState string    `json:"complete_state"`
	SpyCatName    string    `json:"spycat_name"`
	Targets       []string  `json:"targets"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type MissionTargetRequest struct {
	Name          string `json:"name" binding:"required"`
	Country       string `json:"country" binding:"required"`
	Notes         string `json:"notes" binding:"required"`
	CompleteState string `json:"complete_state" binding:"required"`
}

type MissionCreateRequest struct {
	CompleteState string                 `json:"complete_state" binding:"required"`
	Targets       []MissionTargetRequest `json:"targets" binding:"required"`
}

type MissionSingleCreateResponseDto struct {
	Id            uuid.UUID   `json:"id"`
	CompleteState string      `json:"complete_state"`
	Targets       []TargetDto `json:"targets"`
	UpdatedAt     time.Time   `json:"updated_at"`
}

type AssignCatRequest struct {
	SpyCatId uuid.UUID `json:"spycat_id" binding:"required"`
}

type CompletionStateRequest struct {
	CompleteState models.CompleteState `json:"complete_state" binding:"required"`
}

type TargetUpdateRequest struct {
	Notes         string `json:"notes" binding:"omitempty"`
	CompleteState string `json:"complete_state" binding:"omitempty"`
}
