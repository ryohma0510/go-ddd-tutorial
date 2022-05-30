package repository

import "go-ddd-tutorial/domain/model"

type IUserRepository interface {
	FindByID(id model.UserID) (model.User, error)
}
