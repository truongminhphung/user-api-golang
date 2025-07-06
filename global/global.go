package global

import (
	"user-api/pkg/logger"
	"user-api/pkg/setting"
)

// Global configuration variable
var (
	Config setting.Config
	Logger *logger.LoggerZap
)
