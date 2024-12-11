package usecase

import (
	"BookMeHotel/internal/entity/application"
	"BookMeHotel/internal/repo"
	"errors"
	"fmt"
)

type ApplicationUsecase struct {
	applicationRepo repo.ApplicationRepository
	roomRepo        repo.RoomRepository
}

func NewApplicationUsecase(applicationRepo repo.ApplicationRepository, roomRepo repo.RoomRepository) *ApplicationUsecase {
	return &ApplicationUsecase{
		applicationRepo: applicationRepo,
		roomRepo:        roomRepo,
	}
}

func (u *ApplicationUsecase) CreateApplication(app application.Application) error {
	if app.FullName == "" || app.Email == "" || app.Phone == "" {
		return errors.New("all fields are required")
	}

	if app.RoomID <= 0 {
		return errors.New("invalid room ID")
	}

	room, err := u.roomRepo.GetRoomByID(app.RoomID)
	if err != nil {
		return fmt.Errorf("failed to get room: %w", err)
	}

	if room.Capacity < app.GuestCount {
		return errors.New("room capacity exceeded")
	}

	filter := application.Filter{
		RoomID:    &app.RoomID,
		Status:    strPtr("confirmed"),
		StartDate: &app.StartDate,
		EndDate:   &app.EndDate,
	}
	conflictingApps, err := u.applicationRepo.GetApplications(filter)
	if err != nil {
		return fmt.Errorf("failed to retrieve applications: %w", err)
	}

	if len(conflictingApps) > 0 {
		return errors.New("room is already booked for the selected dates")
	}

	return u.applicationRepo.CreateApplication(app)
}

func (u *ApplicationUsecase) DeleteApplication(applicationID int) error {
	if applicationID <= 0 {
		return errors.New("invalid application ID")
	}

	err := u.applicationRepo.DeleteApplication(applicationID)
	if err != nil {
		return fmt.Errorf("failed to delete application: %w", err)
	}

	return nil
}

func (u *ApplicationUsecase) GetApplicationByID(applicationID int) (application.Application, error) {
	if applicationID <= 0 {
		return application.Application{}, errors.New("invalid application ID")
	}

	app, err := u.applicationRepo.GetApplicationByID(applicationID)
	if err != nil {
		return application.Application{}, fmt.Errorf("failed to get application: %w", err)
	}

	return app, nil
}

func (u *ApplicationUsecase) GetApplications(filter application.Filter) ([]application.Application, error) {
	apps, err := u.applicationRepo.GetApplications(filter)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve applications: %w", err)
	}
	return apps, nil
}

func (u *ApplicationUsecase) AddApplicationStatus(applicationID int, status application.Status) (application.Status, error) {
	if applicationID <= 0 {
		return application.Status{}, errors.New("invalid application ID")
	}

	_, err := u.applicationRepo.GetApplicationByID(applicationID)
	if err != nil {
		return application.Status{}, fmt.Errorf("application not found: %w", err)
	}

	err = u.applicationRepo.AddApplicationStatus(applicationID, status)
	if err != nil {
		return application.Status{}, fmt.Errorf("failed to add status: %w", err)
	}

	latestStatus, err := u.applicationRepo.GetLatestStatus(applicationID)
	if err != nil {
		return application.Status{}, fmt.Errorf("failed to retrieve latest status: %w", err)
	}

	return latestStatus, nil
}

func (u *ApplicationUsecase) DeleteApplicationStatus(applicationID int, statusID int) error {
	if applicationID <= 0 || statusID <= 0 {
		return errors.New("invalid application ID or status ID")
	}

	_, err := u.applicationRepo.GetApplicationByID(applicationID)
	if err != nil {
		return fmt.Errorf("application not found: %w", err)
	}

	err = u.applicationRepo.DeleteApplicationStatus(applicationID, statusID)
	if err != nil {
		return fmt.Errorf("failed to delete status: %w", err)
	}

	return nil
}

func (u *ApplicationUsecase) GetAllApplicationStatuses(applicationID int) ([]application.Status, error) {
	if applicationID <= 0 {
		return nil, errors.New("invalid application ID")
	}

	_, err := u.applicationRepo.GetApplicationByID(applicationID)
	if err != nil {
		return nil, fmt.Errorf("application not found: %w", err)
	}

	statuses, err := u.applicationRepo.GetAllApplicationStatuses(applicationID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve statuses: %w", err)
	}

	return statuses, nil
}

func (u *ApplicationUsecase) GetLatestApplicationStatus(applicationID int) (application.Status, error) {
	if applicationID <= 0 {
		return application.Status{}, errors.New("invalid application ID")
	}

	_, err := u.applicationRepo.GetApplicationByID(applicationID)
	if err != nil {
		return application.Status{}, fmt.Errorf("application not found: %w", err)
	}

	status, err := u.applicationRepo.GetLatestStatus(applicationID)
	if err != nil {
		return application.Status{}, fmt.Errorf("failed to retrieve latest status: %w", err)
	}

	return status, nil
}

func strPtr(s string) *string {
	return &s
}
