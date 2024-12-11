package usecase

import (
	"BookMeHotel/internal/entity"
	"BookMeHotel/internal/repo"
	"errors"
	"sync"
)

type RoomUsecase struct {
	roomRepo repo.RoomRepository
}

func NewRoomUsecase(roomRepo repo.RoomRepository) *RoomUsecase {
	return &RoomUsecase{
		roomRepo: roomRepo,
	}
}

func (u *RoomUsecase) CreateRoom(room entity.Room) error {
	if room.RoomNumber != "" || room.Capacity < 0 || room.BasePrice <= 0 {
		//TO:DO
		return errors.New("temp error")
	}
	return u.roomRepo.CreateRoom(room)
}

func (u *RoomUsecase) DeleteRoom(roomID int) error {
	return u.roomRepo.DeleteRoom(roomID)
}

func (u *RoomUsecase) UpdateRoomPrice(roomID int, newPrice float64) error {
	if newPrice < 0 {
		return errors.New("price must be zero or greater")
	}
	return u.roomRepo.UpdateRoomPrice(roomID, newPrice)
}

func (u *RoomUsecase) UpdateMultipleRoomPrices(roomIDs []int, newPrice float64) error {
	var wg sync.WaitGroup
	errChan := make(chan error, len(roomIDs))

	for _, roomID := range roomIDs {
		wg.Add(1)
		go func(roomID int) {
			defer wg.Done()
			err := u.UpdateRoomPrice(roomID, newPrice)
			if err != nil {
				errChan <- err
			}
		}(roomID)
	}
	wg.Wait()
	close(errChan)
	for err := range errChan {
		return err
	}

	return nil
}

func (u *RoomUsecase) GetRoomByID(roomID int) (entity.Room, error) {
	return u.roomRepo.GetRoomByID(roomID)
}

func (u *RoomUsecase) GetAllRooms() ([]entity.Room, error) {
	return u.roomRepo.GetAllRooms()
}
