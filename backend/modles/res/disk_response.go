package res

type GetDisksListResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Capacity     string `json:"capacity"`
	StoragePath  string `json:"storagePath"`
	BelongDevice string `json:"belongDevice"`
	CreatedTime  string `json:"createdTime"`
}
