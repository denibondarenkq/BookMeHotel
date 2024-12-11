package usecase

import (
	"BookMeHotel/internal/entity"
	"BookMeHotel/internal/entity/user"
	"BookMeHotel/internal/repo"
	"errors"
)

type AdminUsecase struct {
	userRepo        repo.UserRepository
	roomRepo        repo.RoomRepository
	applicationRepo repo.ApplicationRepository
}

func NewAdminUsecase(userRepo repo.UserRepository, roomRepo repo.RoomRepository, applicationRepo repo.ApplicationRepository) *AdminUsecase {
	return &AdminUsecase{
		userRepo:        userRepo,
		roomRepo:        roomRepo,
		applicationRepo: applicationRepo,
	}
}

func (u *AdminUsecase) CreateReceptionist(receptionist user.User) error {
	if receptionist.Role != "receptionist" {
		return errors.New("user must have role 'receptionist'")
	}
	return u.userRepo.CreateUser(receptionist)
}

func (u *AdminUsecase) CreateRoom(room entity.Room) error {
	if room.RoomNumber == "" || room.Capacity <= 0 || room.BasePrice <= 0 {
		return errors.New("invalid room data")
	}
	return u.roomRepo.CreateRoom(room)
}
