package application

import "time"

type Status struct {
	ID             int       `json:"id"`
	Status         string    `json:"status"`
	Comment        string    `json:"comment"`
	ReceptionistID int       `json:"receptionist_id"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
}
