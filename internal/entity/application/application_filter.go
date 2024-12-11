package application

import "time"

type Filter struct {
	RoomID     *int
	GuestCount *int
	StartDate  *time.Time
	EndDate    *time.Time
	FoodPlanID *int
	Status     *string
	TotalPrice *float64
}
