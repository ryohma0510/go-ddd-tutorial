package repository

import "go-ddd-tutorial/domain/model"

type ICircleRepository interface {
	Save(circle model.Circle) error
	FindByID(id model.CircleID) (model.Circle, bool, error)
	FindByName(id model.CircleName) (model.Circle, bool, error)
}
