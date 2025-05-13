package models


type DiskBind struct {
	ID       string `json:"id"`
	DeviceId string `json:"device_id"`
	DiskId   string `json:"disk_id"`
}
