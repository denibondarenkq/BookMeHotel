package repo

import (
	"BookMeHotel/internal/entity/application"
)

type ApplicationRepository interface {
	CreateApplication(application application.Application) error
	DeleteApplication(applicationID int) error
	GetApplicationByID(applicationID int) (application.Application, error)
	GetApplications(filter application.Filter) ([]application.Application, error)
	AddApplicationStatus(applicationID int, status application.Status) error
	DeleteApplicationStatus(applicationID int, statusID int) error
	GetLatestStatus(applicationID int) (application.Status, error)
	GetAllApplicationStatuses(applicationID int) ([]application.Status, error)
}
