package services

import (
	"Spy-Cat-Agency/src/internal/spycats/application"
	"Spy-Cat-Agency/src/internal/spycats/domain"
)

type spyCatServiceImpl struct {
	reader  domain.SpyCatReader
	writer  domain.SpyCatWriter
	updater domain.SpyCatUpdater
	deleter domain.SpyCatDeleter
}

func NewSpyCatService(newReader domain.SpyCatReader, newWriter domain.SpyCatWriter, newUpdater domain.SpyCatUpdater, newDeleter domain.SpyCatDeleter) application.SpyCatService {
	return &spyCatServiceImpl{
		reader:  newReader,
		writer:  newWriter,
		updater: newUpdater,
		deleter: newDeleter,
	}
}
