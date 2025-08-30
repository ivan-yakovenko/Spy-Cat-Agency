package application

import (
	"Spy-Cat-Agency/src/internal/spycats/dtos"
	"context"

	"github.com/google/uuid"
)

type SpyCatService interface {
	GetSingleSpyCat(ctx context.Context, id uuid.UUID) (dtos.SpyCatSingleResponseDto, error)
	GetAllSpyCats(ctx context.Context) ([]dtos.SpyCatAllResponseDto, error)

	CreateSpyCat(ctx context.Context, spycatReq dtos.SpyCatCreateRequest) (dtos.SpyCatSingleResponseDto, error)

	UpdateSpyCatSalary(ctx context.Context, salaryReq dtos.SalaryRequest, id uuid.UUID) (dtos.SpyCatSingleResponseDto, error)

	DeleteSingleSpyCat(ctx context.Context, id uuid.UUID) error
	DeleteManySpyCats(ctx context.Context, ids dtos.DeletedIds) error
}
