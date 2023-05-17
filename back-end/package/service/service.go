package service

import (
	"sync"
	"test_puzzle/package/logging"
	"test_puzzle/package/repository"
)

type Service struct {
	Logger     logging.Logger
	mutex      *sync.Mutex
	repository *repository.Repository
}

func New(repository *repository.Repository, mutex *sync.Mutex) *Service {
	return &Service{
		Logger:     logging.GetLogger(),
		mutex:      mutex,
		repository: repository,
	}
}
