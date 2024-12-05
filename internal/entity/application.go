package entity

import "time"

type Application struct {
	ID             int       `json:"id"`
	FullName       string    `json:"full_name"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	RoomID         int       `json:"room_id"`
	GuestCount     int       `json:"guest_count"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	FoodPlanID     int       `json:"food_plan_id"`
	Status         string    `json:"status"`
	ReceptionistID int       `json:"receptionist_id"`
	TotalPrice     float64   `json:"total_price"`
}
