package api

import (
	"gin-vue/api/devices_api"
	"gin-vue/api/disk_api"
	"gin-vue/api/gpu_api"
	"gin-vue/api/user_api"
)

type ApiGroup struct {
	UserApi    user_api.UsersApi
	DevicesApi devices_api.DevicesApi
	DiskApi    disk_api.DisksApi
	GpuApi     gpu_api.GpusApi
}

var ApiGroupApp = new(ApiGroup)
