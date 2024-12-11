package usecase

import (
	"BookMeHotel/internal/entity/user"
	"BookMeHotel/internal/repo"
	"errors"
)

type UserUsecase struct {
	userRepo repo.UserRepository
}

func NewUserUsecase(userRepo repo.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

func (u *UserUsecase) CreateUser(user user.User) error {
	if user.Name == "" || user.Email == "" || user.Password == "" || user.Role == "" {
		return errors.New("all fields are required")
	}

	if user.Role != "admin" && user.Role != "receptionist" {
		return errors.New("invalid role")
	}

	return u.userRepo.CreateUser(user)
}

func (u *UserUsecase) DeleteUser(userID int) error {
	return u.userRepo.DeleteUser(userID)
}

func (u *UserUsecase) GetUserByID(userID int) (user.User, error) {
	return u.userRepo.GetUserByID(userID)
}

func (u *UserUsecase) GetAllUsers() ([]user.User, error) {
	return u.userRepo.GetAllUsers()
}

func (u *UserUsecase) UpdateUserRole(userID int, role string) error {
	return u.userRepo.UpdateUserRole(userID, role)
}
