package service

import (
	"block_chain/repository"

	"github.com/inconshreveable/log15"
)

type Service struct {
	repository *repository.Repository

	log log15.Logger

	difficulty int64
}

func NewService(repository *repository.Repository, difficulty int64) *Service {
	s := &Service{
		repository: repository,
		log:        log15.New("module", "service"),
		difficulty: difficulty,
	}

	return s
}
