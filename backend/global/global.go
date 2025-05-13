package global

import (
	"gin-vue/config"
	"log"

	"gorm.io/gorm"
)


var (
	Config *config.Config
	DB     *gorm.DB
	Logger *log.Logger
)
