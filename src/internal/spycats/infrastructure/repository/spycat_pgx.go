package repository

import (
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"Spy-Cat-Agency/src/internal/spycats/domain/models"
	"context"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SpyCatPgxRepository struct {
	pool *pgxpool.Pool
}

func NewSpyCatPgxRepository(newPool *pgxpool.Pool) *SpyCatPgxRepository {
	return &SpyCatPgxRepository{pool: newPool}
}

func (spr *SpyCatPgxRepository) FindAll(ctx context.Context) ([]models.SpyCat, error) {

	query := `SELECT id, name, experience_years, breed, salary, created_at, updated_at FROM spy_cats`

	rows, err := spr.pool.Query(ctx, query)

	if err != nil {
		return nil, error_handler.NewCustomError(http.StatusInternalServerError, "Error executing spy cats SELECT all query", err)
	}

	defer rows.Close()

	var spycats []models.SpyCat

	for rows.Next() {
		var spycat models.SpyCat
		if err := rows.Scan(
			&spycat.Id,
			&spycat.Name,
			&spycat.ExperienceYears,
			&spycat.Breed,
			&spycat.Salary,
			&spycat.CreatedAt,
			&spycat.UpdatedAt,
		); err != nil {
			return nil, error_handler.NewCustomError(http.StatusInternalServerError, "Error scanning spy cats rows from database", err)
		}
		spycats = append(spycats, spycat)
	}

	if rows.Err() != nil {
		return nil, error_handler.NewCustomError(http.StatusInternalServerError, "Database error", rows.Err())
	}

	return spycats, nil

}

func (spr *SpyCatPgxRepository) FindById(ctx context.Context, id uuid.UUID) (*models.SpyCat, error) {

	var spycat models.SpyCat

	query := `SELECT id, name, experience_years, breed, salary, created_at, updated_at FROM spy_cats WHERE id = $1`

	if err := spr.pool.QueryRow(ctx, query, id).Scan(
		&spycat.Id,
		&spycat.Name,
		&spycat.ExperienceYears,
		&spycat.Breed,
		&spycat.Salary,
		&spycat.CreatedAt,
		&spycat.UpdatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, error_handler.NewCustomError(http.StatusNotFound, "No such spy cat found in the database", err)
		}
		return nil, error_handler.NewCustomError(http.StatusInternalServerError, "Error executing spy cat SELECT query", err)
	}

	return &spycat, nil
}

func (spr *SpyCatPgxRepository) Create(ctx context.Context, newSpyCat *models.SpyCat) (*models.SpyCat, error) {

	query := `INSERT INTO spy_cats (id, name, experience_years, breed, salary) 
			  VALUES ($1, $2, $3, $4, $5)
			  RETURNING id, name, experience_years, breed, salary, created_at, updated_at`

	if err := spr.pool.QueryRow(ctx, query,
		newSpyCat.Id,
		newSpyCat.Name,
		newSpyCat.ExperienceYears,
		newSpyCat.Breed,
		newSpyCat.Salary,
	).Scan(
		&newSpyCat.Id,
		&newSpyCat.Name,
		&newSpyCat.ExperienceYears,
		&newSpyCat.Breed,
		&newSpyCat.Salary,
		&newSpyCat.CreatedAt,
		&newSpyCat.UpdatedAt); err != nil {
		return nil, error_handler.NewCustomError(http.StatusInternalServerError, "Error inserting new spy cat", err)
	}

	return newSpyCat, nil
}

func (spr *SpyCatPgxRepository) Update(ctx context.Context, updatedSpyCat *models.SpyCat) (*models.SpyCat, error) {

	query := `UPDATE spy_cats SET salary = $1, updated_at = NOW() WHERE id = $2 
              RETURNING id, name, experience_years, breed, salary, created_at, updated_at`

	if err := spr.pool.QueryRow(ctx, query,
		updatedSpyCat.Salary,
		updatedSpyCat.Id,
	).Scan(
		&updatedSpyCat.Id,
		&updatedSpyCat.Name,
		&updatedSpyCat.ExperienceYears,
		&updatedSpyCat.Breed,
		&updatedSpyCat.Salary,
		&updatedSpyCat.CreatedAt,
		&updatedSpyCat.UpdatedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, error_handler.NewCustomError(http.StatusInternalServerError, "No spy cats rows to update from database", pgx.ErrNoRows)
		}
		return nil, error_handler.NewCustomError(http.StatusInternalServerError, "Error scanning spy cats rows from database", err)
	}

	return updatedSpyCat, nil
}

func (spr *SpyCatPgxRepository) DeleteById(ctx context.Context, id uuid.UUID) error {

	query := `DELETE FROM spy_cats WHERE id = $1`

	cmdTag, err := spr.pool.Exec(ctx, query, id)

	if err != nil {
		return error_handler.NewCustomError(http.StatusInternalServerError, "Error executing database spy cat DELETE query", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return error_handler.NewCustomError(http.StatusInternalServerError, "No spy cats rows to delete from database", err)
	}

	return nil

}

func (spr *SpyCatPgxRepository) DeleteMany(ctx context.Context, ids []uuid.UUID) error {

	query := `DELETE FROM spy_cats WHERE id = ANY($1)`

	cmdTag, err := spr.pool.Exec(ctx, query, ids)

	if err != nil {
		return error_handler.NewCustomError(http.StatusInternalServerError, "Error executing database spy cats DELETE query", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return error_handler.NewCustomError(http.StatusInternalServerError, "No spy cat rows to delete from database", err)
	}

	return nil

}
