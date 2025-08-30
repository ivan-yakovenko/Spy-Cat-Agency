package repository

import (
	"Spy-Cat-Agency/src/internal/missions/domain/models"
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	"context"

	"github.com/go-faster/errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MissionPgxRepository struct {
	pool *pgxpool.Pool
}

func NewMissionPgxRepository(newPool *pgxpool.Pool) *MissionPgxRepository {
	return &MissionPgxRepository{pool: newPool}
}

func (spr *MissionPgxRepository) FindAllMissions(ctx context.Context) ([]models.Mission, error) {

	query := `SELECT id, spycat_id, complete_state, created_at, updated_at FROM missions`

	rows, err := spr.pool.Query(ctx, query)

	if err != nil {
		return nil, error_handler.ErrorHandler(err, err.Error())
	}

	defer rows.Close()

	var missions []models.Mission

	for rows.Next() {
		var mission models.Mission
		if err := rows.Scan(
			&mission.Id,
			&mission.SpyCatId,
			&mission.CompleteState,
			&mission.CreatedAt,
			&mission.UpdatedAt,
		); err != nil {
			return nil, error_handler.ErrorHandler(err, err.Error())
		}
		missions = append(missions, mission)
	}

	if rows.Err() != nil {
		return nil, error_handler.ErrorHandler(err, rows.Err().Error())
	}

	return missions, nil

}

func (spr *MissionPgxRepository) FindMissionById(ctx context.Context, id uuid.UUID) (*models.Mission, error) {

	var mission models.Mission

	query := `SELECT id, spycat_id, complete_state, created_at, updated_at FROM missions WHERE id = $1`

	if err := spr.pool.QueryRow(ctx, query, id).Scan(
		&mission.Id,
		&mission.SpyCatId,
		&mission.CompleteState,
		&mission.CreatedAt,
		&mission.UpdatedAt,
	); err != nil {
		return nil, error_handler.ErrorHandler(err, err.Error())
	}

	return &mission, nil
}

func (spr *MissionPgxRepository) FindMissionTargetsById(ctx context.Context, missionId uuid.UUID) ([]models.Target, error) {

	query := `SELECT id, mission_id, name, country, notes, complete_state, created_at, updated_at FROM targets WHERE mission_id = $1`

	rows, err := spr.pool.Query(ctx, query, missionId)

	if err != nil {
		return nil, error_handler.ErrorHandler(err, err.Error())
	}

	defer rows.Close()

	var targets []models.Target

	for rows.Next() {
		var target models.Target
		if err := rows.Scan(
			&target.Id,
			&target.MissionId,
			&target.Name,
			&target.Country,
			&target.Notes,
			&target.CompleteState,
			&target.CreatedAt,
			&target.UpdatedAt,
		); err != nil {
			return nil, error_handler.ErrorHandler(err, err.Error())
		}
		targets = append(targets, target)
	}

	if rows.Err() != nil {
		return nil, error_handler.ErrorHandler(err, rows.Err().Error())
	}

	return targets, nil

}

func (spr *MissionPgxRepository) FindTargetById(ctx context.Context, targetId uuid.UUID) (*models.Target, error) {

	var target models.Target

	query := `SELECT id, mission_id, name, country, notes, complete_state, created_at, updated_at FROM targets WHERE id = $1`

	if err := spr.pool.QueryRow(ctx, query, targetId).Scan(
		&target.Id,
		&target.MissionId,
		&target.Name,
		&target.Country,
		&target.Notes,
		&target.CompleteState,
		&target.CreatedAt,
		&target.UpdatedAt,
	); err != nil {
		return nil, error_handler.ErrorHandler(err, err.Error())
	}

	return &target, nil

}

func (spr *MissionPgxRepository) Create(ctx context.Context, tx pgx.Tx, newMission *models.Mission, newTargets []models.Target) (*models.Mission, error) {

	newMissionQuery := `INSERT INTO missions (id, spycat_id, complete_state) VALUES ($1, $2, $3) RETURNING id, complete_state, created_at, updated_at`

	if err := tx.QueryRow(ctx, newMissionQuery, newMission.Id, newMission.SpyCatId,
		newMission.CompleteState).Scan(
		&newMission.Id,
		&newMission.CompleteState,
		&newMission.CreatedAt,
		&newMission.UpdatedAt); err != nil {
		return nil, error_handler.ErrorHandler(err, err.Error())
	}

	newTargetQuery := `INSERT INTO targets (id, mission_id, name, country, notes, complete_state) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, name, country, notes, complete_state, created_at, updated_at`

	for _, target := range newTargets {

		if err := tx.QueryRow(ctx, newTargetQuery, target.Id, target.MissionId,
			target.Name,
			target.Country,
			target.Notes,
			target.CompleteState).Scan(
			&target.Id,
			&target.Name,
			&target.Country,
			&target.Notes,
			&target.CompleteState,
			&target.CreatedAt,
			&target.UpdatedAt,
		); err != nil {
			return nil, error_handler.ErrorHandler(err, err.Error())
		}
	}

	return newMission, nil
}

func (spr *MissionPgxRepository) AssignCatToMission(ctx context.Context, tx pgx.Tx, mission *models.Mission) (*models.Mission, error) {

	query := `UPDATE missions SET spycat_id = $1 WHERE id = $2 
              RETURNING id, spycat_id, complete_state, created_at, updated_at`

	if err := tx.QueryRow(ctx, query, mission.SpyCatId, mission.Id).Scan(
		&mission.Id,
		&mission.SpyCatId,
		&mission.CompleteState,
		&mission.CreatedAt,
		&mission.UpdatedAt); err != nil {
		return nil, error_handler.ErrorHandler(err, err.Error())
	}

	return mission, nil
}

func (spr *MissionPgxRepository) CreateTarget(ctx context.Context, tx pgx.Tx, target *models.Target) (*models.Target, error) {
	query := `INSERT INTO targets (id, mission_id, name, country, notes, complete_state)
			  VALUES ($1, $2, $3, $4, $5, $6)
			  RETURNING id, mission_id, name, country, notes, complete_state, created_at, updated_at`

	var createdTarget models.Target

	if err := tx.QueryRow(ctx, query, target.Id, target.MissionId, target.Name,
		target.Country,
		target.Notes,
		target.CompleteState).Scan(
		&createdTarget.Id,
		&createdTarget.MissionId,
		&createdTarget.Name,
		&createdTarget.Country,
		&createdTarget.Notes,
		&createdTarget.CompleteState,
		&createdTarget.CreatedAt,
		&createdTarget.UpdatedAt); err != nil {
		return nil, error_handler.ErrorHandler(err, err.Error())
	}

	return &createdTarget, nil
}

func (spr *MissionPgxRepository) UpdateMission(ctx context.Context, tx pgx.Tx, updatedMission *models.Mission) (*models.Mission, error) {

	query := `UPDATE missions SET complete_state = $1 WHERE id = $2 
              RETURNING id, spycat_id, complete_state, created_at, updated_at`

	if err := tx.QueryRow(ctx, query,
		updatedMission.CompleteState,
		updatedMission.Id,
	).Scan(
		&updatedMission.Id,
		&updatedMission.SpyCatId,
		&updatedMission.CompleteState,
		&updatedMission.CreatedAt,
		&updatedMission.UpdatedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, error_handler.ErrorHandler(err, pgx.ErrNoRows.Error())
		}
		return nil, error_handler.ErrorHandler(err, err.Error())
	}

	return updatedMission, nil
}

func (spr *MissionPgxRepository) UpdateTarget(ctx context.Context, tx pgx.Tx, updatedTarget *models.Target) (*models.Target, error) {

	query := `UPDATE targets SET notes = $1, complete_state = $2 WHERE id = $3
			  RETURNING id, mission_id, name, country, notes, 
			      complete_state, created_at, updated_at`

	if err := tx.QueryRow(ctx, query, updatedTarget.Notes, updatedTarget.CompleteState,
		updatedTarget.Id).Scan(
		&updatedTarget.Id,
		&updatedTarget.MissionId,
		&updatedTarget.Name,
		&updatedTarget.Country,
		&updatedTarget.Notes,
		&updatedTarget.CompleteState,
		&updatedTarget.CreatedAt,
		&updatedTarget.UpdatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, error_handler.ErrorHandler(err, pgx.ErrNoRows.Error())
		}
		return nil, error_handler.ErrorHandler(err, err.Error())
	}

	return updatedTarget, nil
}

func (spr *MissionPgxRepository) DeleteMissionById(ctx context.Context, tx pgx.Tx, id uuid.UUID) error {

	query := `DELETE FROM missions WHERE id = $1`

	cmdTag, err := tx.Exec(ctx, query, id)

	if err != nil {
		return error_handler.ErrorHandler(err, err.Error())
	}

	if cmdTag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil

}

func (spr *MissionPgxRepository) DeleteTargetById(ctx context.Context, tx pgx.Tx, id uuid.UUID) error {

	query := `DELETE FROM targets WHERE id = $1`

	cmdTag, err := tx.Exec(ctx, query, id)

	if err != nil {
		return error_handler.ErrorHandler(err, err.Error())
	}

	if cmdTag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil

}
