package initialize

import (
	"user-api/global"
	"user-api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLoggerZap(global.Config.Logger)
}
