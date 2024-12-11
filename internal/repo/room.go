package repo

import "BookMeHotel/internal/entity"

type RoomRepository interface {
	CreateRoom(room entity.Room) error
	DeleteRoom(roomID int) error
	GetRoomByID(roomID int) (entity.Room, error)
	GetAllRooms() ([]entity.Room, error)
	UpdateRoomPrice(roomID int, newPrice float64) error
}
