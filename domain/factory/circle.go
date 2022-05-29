package factory

import "go-ddd-tutorial/domain/model"

type ICircleFactory interface {
	Create(name model.CircleName, owner model.User) (model.Circle, error)
}
