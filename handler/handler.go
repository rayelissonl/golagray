package handler

import (
	"gorm.io/gorm"

	"github.com/rayelissonl/goapreapi/config"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
