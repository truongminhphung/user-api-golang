package user

import (
	"log"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router
	log.Println("Initializing User Routes")
	userRouterPublic := Router.Group("/user")
	{
		// Update this endpoint later
		userRouterPublic.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  "ok",
				"message": "List all users",
			})
		})
		userRouterPublic.POST("", func(c *gin.Context) {
			c.JSON(201, gin.H{
				"status":  "created",
				"message": "User created successfully",
			})
		})
		userRouterPublic.GET("/:id", func(c *gin.Context) {
			userID := c.Param("id")
			c.JSON(200, gin.H{
				"status":  "ok",
				"message": "User details",
				"user_id": userID,
			})
		})
		userRouterPublic.PUT("/:id", func(c *gin.Context) {
			userID := c.Param("id")
			c.JSON(200, gin.H{
				"status":  "ok",
				"message": "User updated successfully",
				"user_id": userID,
			})
		})
		userRouterPublic.DELETE("/:id", func(c *gin.Context) {
			userID := c.Param("id")
			c.JSON(200, gin.H{
				"status":  "ok",
				"message": "User deleted successfully",
				"user_id": userID,
			})
		})
	}

}
