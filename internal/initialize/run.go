package initialize

import "github.com/gin-gonic/gin"

func Run() *gin.Engine {
	// Load configuration
	LoadConfig()

	r := InitRouter()

	return r
}
