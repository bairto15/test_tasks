package service

import (
	"test_puzzle/package/logging"
	"test_puzzle/package/repository"
)

type Service struct {
	Logger     logging.Logger
	repository *repository.Repository
}

func New(repository *repository.Repository) *Service {
	return &Service{
		Logger: logging.GetLogger(),
		repository: repository,
	}
}
