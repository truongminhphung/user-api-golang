package initialize

import "github.com/gin-gonic/gin"

func Run() *gin.Engine {
	// Load configuration
	LoadConfig()

	return gin.Default()
}
