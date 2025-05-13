package models


type License struct {
	ID          string `json:"id"`
	LicenseCode string `json:"license_code"`
	LicenseType string `json:"license_type"`
	ActiveTime  int64  `json:"active_time"`
	ExpiredTime int64  `json:"expired_time"`
	ExpireFlag  int64  `json:"expire_flag"`
}
