package repo

import (
	"BookMeHotel/internal/entity/user"
)

type UserRepository interface {
	CreateUser(user user.User) error
	DeleteUser(userID int) error
	GetUserByID(userID int) (user.User, error)
	GetAllUsers() ([]user.User, error)
	UpdateUserRole(userID int, role string) error
}
