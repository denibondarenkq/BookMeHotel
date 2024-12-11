package user

type Receptionist struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
	AdminID  int    `json:"admin_id"`
}
