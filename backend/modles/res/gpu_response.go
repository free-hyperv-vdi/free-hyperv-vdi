package res

type GetGpusListResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	InstancePath string `json:"instancePath"`
	BindCount    int    `json:"bindCount"`
}
