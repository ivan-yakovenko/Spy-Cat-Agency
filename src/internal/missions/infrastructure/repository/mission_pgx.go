package repository

import (
	"Spy-Cat-Agency/src/internal/missions/domain/models"
	"Spy-Cat-Agency/src/internal/shared/utils/error_handler"
	models2 "Spy-Cat-Agency/src/internal/spycats/domain/models"
	"context"
	"net/http"

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

func (mpr *MissionPgxRepository) FindAllMissions(ctx context.Context) (map[uuid.UUID]*models.MissionDetails, error) {

	query := `SELECT m.id, m.spycat_id, m.complete_state, m.created_at, m.updated_at,
						sc.name AS spycat_name,
						t.name AS target_name
			FROM missions m
					 LEFT JOIN spy_cats sc ON m.spycat_id = sc.id
					 LEFT JOIN targets t ON m.id = t.mission_id`

	rows, err := mpr.pool.Query(ctx, query)

	if err != nil {
		return nil, error_handler.NewCustomError(http.StatusInternalServerError, "Error executing mission SELECT all query", err)
	}

	missionsData := make(map[uuid.UUID]*models.MissionDetails)

	for rows.Next() {
		var (
			mission    models.Mission
			targetName *string
			spyCatName *string
		)
		if err := rows.Scan(
			&mission.Id,
			&mission.SpyCatId,
			&mission.CompleteState,
			&mission.CreatedAt,
			&mission.UpdatedAt,
			&spyCatName,
			&targetName,
		); err != nil {
			return nil, error_handler.NewCustomError(http.StatusInternalServerError, "Error scanning missions rows from database", err)
		}

		m, ok := missionsData[mission.Id]

		if !ok {
			m = &models.MissionDetails{
				Mission:     mission,
				SpyCatName:  spyCatName,
				TargetNames: []string{},
			}
			missionsData[mission.Id] = m
		}

		if targetName != nil {
			m.TargetNames = append(m.TargetNames, *targetName)
		}
	}

	if rows.Err() != nil {
		return nil, error_handler.NewCustomError(http.StatusInternalServerError, "Database error", rows.Err())
	}

	return missionsData, nil

}

func (mpr *MissionPgxRepository) FindMissionById(ctx context.Context, id uuid.UUID) (*models.Mission, []models.Target, *models2.SpyCat, error) {

	query := `SELECT m.id, m.spycat_id, m.complete_state, m.created_at, m.updated_at,
       				 sc.id, 
       				 coalesce(sc.name, '') as name,
       				 coalesce(sc.experience_years, 0) as experience_years, 
       				 coalesce(sc.breed, '') as breed,
       				 coalesce(sc.salary, 0) as salary,
       				 coalesce(sc.created_at, '1970-01-01 00:00:00') as created_at,
       				 coalesce (sc.updated_at, '1970-01-01 00:00:00') as updated_at,
       				 t.id, t.mission_id, t.name, t.country, t.notes, t.complete_state, t.created_at, t.updated_at
       FROM missions m 
       LEFT JOIN spy_cats sc ON m.spycat_id = sc.id
       LEFT JOIN targets t ON m.id = t.mission_id
       WHERE m.id = $1`

	rows, err := mpr.pool.Query(ctx, query, id)
	if err != nil {
		return nil, nil, nil, error_handler.NewCustomError(http.StatusInternalServerError, "Error executing mission SELECT query", err)
	}

	defer rows.Close()

	var (
		mission *models.Mission
		spycat  *models2.SpyCat
		targets []models.Target
	)

	for rows.Next() {
		var (
			m    models.Mission
			sc   models2.SpyCat
			t    models.Target
			scId *uuid.UUID
			tId  *uuid.UUID
		)

		if err := rows.Scan(
			&m.Id, &scId, &m.CompleteState, &m.CreatedAt, &m.UpdatedAt,
			&sc.Id, &sc.Name, &sc.ExperienceYears, &sc.Breed, &sc.Salary, &sc.CreatedAt, &sc.UpdatedAt,
			&tId, &t.MissionId, &t.Name, &t.Country, &t.Notes, &t.CompleteState, &t.CreatedAt, &t.UpdatedAt,
		); err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return nil, nil, nil, error_handler.NewCustomError(http.StatusNotFound, "No such mission found in the database", err)
			}
			return nil, nil, nil, error_handler.NewCustomError(http.StatusInternalServerError, "Error scanning missions rows from database", err)
		}
		if mission == nil {
			m.SpyCatId = scId
			mission = &m
			if scId != nil {
				spycat = &sc
			} else {
				spycat = &models2.SpyCat{}
			}
		}
		if tId != nil {
			t.Id = *tId
			targets = append(targets, t)
		}
	}

	if mission == nil {
		return nil, nil, nil, error_handler.NewCustomError(http.StatusInternalServerError, "No missions found in database", err)
	}

	return mission, targets, spycat, nil
}

func (mpr *MissionPgxRepository) FindMissionTargetsById(ctx context.Context, missionId uuid.UUID) ([]models.Target, error) {

	query := `SELECT id, mission_id, name, country, notes, complete_state, created_at, updated_at FROM targets WHERE mission_id = $1`

	rows, err := mpr.pool.Query(ctx, query, missionId)

	if err != nil {
		return nil, error_handler.NewCustomError(http.StatusInternalServerError, "Error executing targets SELECT query", err)
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
			return nil, error_handler.NewCustomError(http.StatusInternalServerError, "Error scanning targets rows from database", err)
		}
		targets = append(targets, target)
	}

	if rows.Err() != nil {
		return nil, error_handler.NewCustomError(http.StatusInternalServerError, "Database error", err)
	}

	return targets, nil

}

func (mpr *MissionPgxRepository) FindTargetById(ctx context.Context, targetId uuid.UUID) (*models.Target, error) {

	var target models.Target

	query := `SELECT id, mission_id, name, country, notes, complete_state, created_at, updated_at FROM targets WHERE id = $1`

	if err := mpr.pool.QueryRow(ctx, query, targetId).Scan(
		&target.Id,
		&target.MissionId,
		&target.Name,
		&target.Country,
		&target.Notes,
		&target.CompleteState,
		&target.CreatedAt,
		&target.UpdatedAt,
	); err != nil {
		return nil, error_handler.NewCustomError(http.StatusInternalServerError, "Error executing target SELECT query", err)
	}

	return &target, nil

}

func (mpr *MissionPgxRepository) Create(ctx context.Context, tx pgx.Tx, newMission *models.Mission, newTargets []models.Target) (*models.Mission, error) {

	newMissionQuery := `INSERT INTO missions (id, spycat_id, complete_state) VALUES ($1, $2, $3) RETURNING id, complete_state, created_at, updated_at`

	if err := tx.QueryRow(ctx, newMissionQuery, newMission.Id, newMission.SpyCatId, newMission.CompleteState).Scan(
		&newMission.Id,
		&newMission.CompleteState,
		&newMission.CreatedAt,
		&newMission.UpdatedAt); err != nil {
		return nil, error_handler.NewCustomError(http.StatusInternalServerError, "Error inserting new mission", err)
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
			return nil, error_handler.NewCustomError(http.StatusInternalServerError, "Error inserting new target", err)
		}
	}

	return newMission, nil
}

func (mpr *MissionPgxRepository) AssignCatToMission(ctx context.Context, tx pgx.Tx, mission *models.Mission) (*models.Mission, error) {

	query := `UPDATE missions SET spycat_id = $1 WHERE id = $2 
              RETURNING id, spycat_id, complete_state, created_at, updated_at`

	if err := tx.QueryRow(ctx, query, mission.SpyCatId, mission.Id).Scan(
		&mission.Id,
		&mission.SpyCatId,
		&mission.CompleteState,
		&mission.CreatedAt,
		&mission.UpdatedAt); err != nil {
		return nil, error_handler.NewCustomError(http.StatusInternalServerError, "Error updating mission with the new spy cat", err)
	}

	return mission, nil
}

func (mpr *MissionPgxRepository) CreateTarget(ctx context.Context, tx pgx.Tx, target *models.Target) (*models.Target, error) {
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
		return nil, error_handler.NewCustomError(http.StatusInternalServerError, "Error inserting new target", err)
	}

	return &createdTarget, nil
}

func (mpr *MissionPgxRepository) UpdateMission(ctx context.Context, tx pgx.Tx, updatedMission *models.Mission) (*models.Mission, error) {

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
			return nil, error_handler.NewCustomError(http.StatusInternalServerError, "No rows to update in missions from database", pgx.ErrNoRows)
		}
		return nil, error_handler.NewCustomError(http.StatusInternalServerError, "Error updating mission", err)
	}

	return updatedMission, nil
}

func (mpr *MissionPgxRepository) UpdateTarget(ctx context.Context, tx pgx.Tx, updatedTarget *models.Target) (*models.Target, error) {

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
			return nil, error_handler.NewCustomError(http.StatusInternalServerError, "No rows to update in targets from database", pgx.ErrNoRows)
		}
		return nil, error_handler.NewCustomError(http.StatusInternalServerError, "Error updating target", err)
	}

	return updatedTarget, nil
}

func (mpr *MissionPgxRepository) DeleteMissionById(ctx context.Context, tx pgx.Tx, id uuid.UUID) error {

	query := `DELETE FROM missions WHERE id = $1`

	cmdTag, err := tx.Exec(ctx, query, id)

	if err != nil {
		return error_handler.NewCustomError(http.StatusInternalServerError, "Error executing database DELETE query for mission", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return error_handler.NewCustomError(http.StatusInternalServerError, "No rows to delete in missions from database", err)
	}

	return nil

}

func (mpr *MissionPgxRepository) DeleteTargetById(ctx context.Context, tx pgx.Tx, id uuid.UUID) error {

	query := `DELETE FROM targets WHERE id = $1`

	cmdTag, err := tx.Exec(ctx, query, id)

	if err != nil {
		return error_handler.NewCustomError(http.StatusInternalServerError, "Error executing database DELETE query for target", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return error_handler.NewCustomError(http.StatusInternalServerError, "No rows to delete in targets from database", err)
	}

	return nil

}
