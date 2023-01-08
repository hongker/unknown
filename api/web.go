package api

type LoginRequest struct {
	Phone string `json:"phone"`
}
type LoginResponse struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}
