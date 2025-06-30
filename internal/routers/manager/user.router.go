package manager

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouterPrivate := Router.Group("/admin/user")
	{
		userRouterPrivate.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  "ok",
				"message": "List all users",
			})
		})
	}
}
