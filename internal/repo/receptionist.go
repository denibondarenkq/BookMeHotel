package repo

import (
	"BookMeHotel/internal/entity/user"
)

type ReceptionistRepository interface {
	CreateReceptionist(receptionist user.Receptionist) error
	DeleteReceptionist(receptionist user.Receptionist) error
	GetAllReceptionists() ([]user.Receptionist, error)
}
