package models


type Disk struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Capacity    string `json:"capacity"`
	StoragePath string `json:"storage_path"`
	CreatedTime int64  `json:"created_time"`
}
