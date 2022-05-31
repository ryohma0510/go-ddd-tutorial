package usecase

import (
	"errors"
	"fmt"
	"go-ddd-tutorial/domain/factory"
	"go-ddd-tutorial/domain/model"
	"go-ddd-tutorial/domain/repository"
	"go-ddd-tutorial/domain/service"
	"go-ddd-tutorial/domain/spec"
)

var (
	errCircleDuplicated = errors.New("circle duplicated")
)

type CircleAppService struct {
	circleFactory       factory.ICircleFactory
	circleRepository    repository.ICircleRepository
	circleDomainService service.CircleDomainService
	userRepository      repository.IUserRepository
}

func NewCircleAppService(
	circleFactory factory.ICircleFactory,
	circleRepository repository.ICircleRepository,
	circleDomainService service.CircleDomainService,
	userRepository repository.IUserRepository,
) CircleAppService {
	return CircleAppService{
		circleFactory:       circleFactory,
		circleRepository:    circleRepository,
		circleDomainService: circleDomainService,
		userRepository:      userRepository,
	}
}

func (service CircleAppService) Create(ownerID string, circleName string) error {
	owner, err := service.userRepository.FindByID(model.NewUserID(ownerID))
	if err != nil {
		return fmt.Errorf("cannot find owner when creating circle: %w", err)
	}

	newCircleName, err := model.NewCircleName(circleName)
	if err != nil {
		return err
	}

	newCircle, err := service.circleFactory.Create(newCircleName, owner)
	if err != nil {
		return fmt.Errorf("cannot create new circle: %w", err)
	}
	exists, err := service.circleDomainService.Exists(newCircle)
	if err != nil {
		return err
	}
	if exists {
		return errCircleDuplicated
	}

	if err := service.circleRepository.Save(newCircle); err != nil {
		return err
	}

	return nil
}

func (service CircleAppService) Join(userID string, circleID string) error {
	joiningUser, err := service.userRepository.FindByID(model.NewUserID(userID))
	if err != nil {
		return err
	}

	joiningCircleID, err := model.NewCircleID(circleID)
	if err != nil {
		return err
	}
	joiningCircle, err := service.circleRepository.FindByID(joiningCircleID)
	if err != nil {
		return err
	}

	circleFullSpec := spec.NewCircleFullSpec(service.userRepository)
	if !circleFullSpec.IsSatisfiedByCircle(joiningCircle) {
		return errors.New("circle is full")
	}

	if err := joiningCircle.Join(joiningUser.ID()); err != nil {
		return err
	}

	if err := service.circleRepository.Save(joiningCircle); err != nil {
		return err
	}

	return nil
}
