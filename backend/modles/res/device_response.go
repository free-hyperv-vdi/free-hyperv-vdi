package res

type GetDevicesListResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	UserName     string `json:"username"`
	VirtualIp    string `json:"virtualIp"`
	TemplateInfo string `json:"templateInfo"`
	Status       string `json:"status"`
	CreatedTime  string `json:"createdTime"`
	MemoryInfo   string `json:"memoryInfo"`
	CpuInfo      string `json:"cpuInfo"`
	GpuInfo      string `json:"gpuInfo"`
}

type GetTemplatesListResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
	UserPwd  string `json:"userpwd"`
}
