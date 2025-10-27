package service

import "encontradev/internal/repository"

type Service struct {
	Repository *repository.Repository
}

func RegisterService(repository *repository.Repository) (service *Service, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	serv := &Service{
		Repository: repository,
	}

	return serv, err
}
