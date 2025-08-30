package domain

import (
	"Spy-Cat-Agency/src/internal/spycats/domain/models"
	"context"

	"github.com/google/uuid"
)

type SpyCatReader interface {
	FindAll(ctx context.Context) ([]models.SpyCat, error)
	FindById(ctx context.Context, id uuid.UUID) (*models.SpyCat, error)
}

type SpyCatWriter interface {
	Create(ctx context.Context, newSpyCat *models.SpyCat) (*models.SpyCat, error)
}

type SpyCatUpdater interface {
	Update(ctx context.Context, updatedSpyCat *models.SpyCat) (*models.SpyCat, error)
}

type SpyCatDeleter interface {
	DeleteById(ctx context.Context, id uuid.UUID) error
	DeleteMany(ctx context.Context, ids []uuid.UUID) error
}
