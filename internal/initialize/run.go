package initialize

import (
	"user-api/global"

	"github.com/gin-gonic/gin"
)

func Run() *gin.Engine {
	// Load configuration
	LoadConfig()
	InitLogger()
	global.Logger.Info("Initialized logger successfully.")
	InitMySQL()

	r := InitRouter()
	global.Logger.Info("Initialized router successfully.")
	return r
}
