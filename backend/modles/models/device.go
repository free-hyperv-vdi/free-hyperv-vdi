package models


type Device struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Ip          string `json:"ip"`
	Status      string `json:"status"`
	CreatedTime int64  `json:"created_time"`
	CpuInfo     string `json:"cpu_info"`
	MemoryInfo  string `json:"memory_info"`
	GpuInfo     string `json:"gpu_info"`
	TemplateId  string `json:"template_id"`
}
