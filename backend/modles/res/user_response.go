package res

type GetUserProfileReponse struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type GetUserLicenseResponse struct {
	MachineCode string `json:"machineCode"`
	IsChecked   bool   `json:"isChecked"`
}

type GetLicensesResponse struct {
	MachineCode string `json:"machineCode"`
	LicenseCode string `json:"licenseCode"`
	LicenseType string `json:"licenseType"`
	ExpireTime  string `json:"expireTime"`
}
