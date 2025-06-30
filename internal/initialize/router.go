package initialize

import (
	"user-api/global"
	"user-api/internal/routers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	managerRouter := routers.RouterGroupApp.Manager
	userRouter := routers.RouterGroupApp.User

	// Add a default route FIRST to ensure the server is working
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status":  "ok",
			"message": "Welcome to User API",
			"port":    global.Config.Server.Port,
			"mode":    global.Config.Server.Mode,
		})
	})

	mainGroup := r.Group("/api/v1")
	{
		mainGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  "ok",
				"message": "Server is running",
			})
		})
	}
	userRouter.InitUserRouter(mainGroup)
	managerRouter.InitUserRouter(mainGroup)
	managerRouter.InitAdminRouter(mainGroup)

	return r
}
