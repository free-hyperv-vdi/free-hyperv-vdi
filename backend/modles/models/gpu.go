package models



type Gpu struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	InstancePath string `json:"instance_path"`
	Desc         string `json:"desc"`
}
