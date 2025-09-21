package models

import (
	"time"

	"github.com/google/uuid"
)

type CompleteState string

var (
	InProgress CompleteState = "in progress"
	Completed  CompleteState = "completed"
)

type Mission struct {
	Id            uuid.UUID     `json:"id"`
	SpyCatId      *uuid.UUID    `json:"spycat_id"`
	CompleteState CompleteState `json:"complete_state"`
	CreatedAt     time.Time
	UpdatedAt     time.Time `json:"updated_at"`
}

type Target struct {
	Id            uuid.UUID     `json:"id"`
	MissionId     uuid.UUID     `json:"mission_id"`
	Name          string        `json:"name"`
	Country       string        `json:"country"`
	Notes         string        `json:"notes"`
	CompleteState CompleteState `json:"complete_state"`
	CreatedAt     time.Time
	UpdatedAt     time.Time `json:"updated_at"`
}

type MissionDetails struct {
	Mission
	SpyCatName  *string
	TargetNames []string
}
