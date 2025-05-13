package models


type Bind struct {
	ID       string `json:"id"`
	DeviceId string `json:"device_id"`
	UserId   string `json:"user_id"`
}
