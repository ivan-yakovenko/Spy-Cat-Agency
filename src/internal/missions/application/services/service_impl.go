package services

import (
	"Spy-Cat-Agency/src/internal/missions/application"
	"Spy-Cat-Agency/src/internal/missions/domain"
	domain2 "Spy-Cat-Agency/src/internal/spycats/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type missionServiceImpl struct {
	reader       domain.MissionReader
	writer       domain.MissionWriter
	updater      domain.MissionUpdater
	deleter      domain.MissionDeleter
	spyCatReader domain2.SpyCatReader
	pool         *pgxpool.Pool
}

func NewMissionService(newReader domain.MissionReader, newWriter domain.MissionWriter,
	newUpdater domain.MissionUpdater, newDeleter domain.MissionDeleter,
	newSpyCatReader domain2.SpyCatReader, newPool *pgxpool.Pool) application.MissionService {
	return &missionServiceImpl{
		reader:       newReader,
		writer:       newWriter,
		updater:      newUpdater,
		deleter:      newDeleter,
		spyCatReader: newSpyCatReader,
		pool:         newPool,
	}
}
