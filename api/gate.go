package api

type GateLoginRequest struct {
	Token string `json:"token"`
}
type GateLoginResponse struct {
	ServerTime int64 `json:"server_time"`
}

type GateHeartbeatRequest struct{}
type GateHeartbeatResponse struct{}
