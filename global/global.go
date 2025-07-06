package global

import (
	"user-api/pkg/logger"
	"user-api/pkg/setting"

	"gorm.io/gorm"
)

// Global configuration variable
var (
	Config  setting.Config    // Config is the global variable for the application configuration
	Logger  *logger.LoggerZap // Logger is the global variable for the logger instance
	MySQLDB *gorm.DB          // MySQLDB is the global variable for the MySQL database connection
)
