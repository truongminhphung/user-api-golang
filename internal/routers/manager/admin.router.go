package manager

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (pr *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	// public router
	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.POST("/login", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"status":  "ok",
				"message": "Admin login successful",
			})
		})
	}
}
