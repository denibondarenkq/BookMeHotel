package entity

type Room struct {
	ID          int     `json:"id"`
	RoomNumber  string  `json:"room_number"`
	Capacity    int     `json:"capacity"`
	Description string  `json:"description"`
	BasePrice   float64 `json:"base_price"`
	AdminID     int     `json:"admin_id"`
}
