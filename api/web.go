package api

type WebLoginRequest struct {
	Phone string `json:"phone"`
}
type WebLoginResponse struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}
