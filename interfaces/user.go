package interfaces

import "PilEk/model"

type IUser interface {
	GetByUsername(string) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(user *model.User) error
	List(offset, limit int) ([]model.User, int, error)
}
