package service

import (
	"go-ddd-tutorial/domain/model"
	"go-ddd-tutorial/domain/repository"
)

type CircleDomainService struct {
	circleRepository repository.ICircleRepository
}

func NewCircleDomainService(repo repository.ICircleRepository) (CircleDomainService, error) {
	return CircleDomainService{circleRepository: repo}, nil
}

func (service CircleDomainService) Exists(circle model.Circle) (bool, error) {
	_, err := service.circleRepository.FindByName(circle.Name())
	if err != nil {
		return false, err
	}

	return true, nil
}
