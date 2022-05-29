package service

import (
	"fmt"
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
	_, isFound, err := service.circleRepository.FindByName(circle.Name())
	if err != nil {
		return false, fmt.Errorf("unknown error while finding circle by name: %w", err)
	}

	return isFound, nil
}
