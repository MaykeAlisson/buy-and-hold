package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/maykealisson/buy-and-hold/controllers"
)

func HandlerRequest() {
	r := gin.Default()

	r.GET("/", controllers.Swagger)

	// Auth
	r.POST("/api/v1/auth", controllers.Auth)

	// User
	r.POST("/api/v1/users", controllers.CreateUser)
	r.PUT("/api/v1/users/:id", controllers.UpdateUser)
	r.DELETE("/api/v1/users/:id", controllers.DeleteUser)

	// Assert
	r.GET("/api/v1/asserts", controllers.GetAssertBy)
	r.POST("/api/v1/asserts", controllers.CreateAssert)
	r.PUT("/api/v1/asserts/:id", controllers.UpdateAssert)
	r.DELETE("/api/v1/asserts/:id", controllers.DeleteAssert)

	// Launch
	r.GET("/api/v1/launchs/:month", controllers.GetByMonth)
	r.GET("/api/v1/:assert/launchs", controllers.GetByAssert)
	r.POST("/api/v1/:assert/launchs", controllers.CreateLaunch)
	r.PUT("/api/v1/:assert/launchs/:id", controllers.UpdateLaunch)
	r.DELETE("/api/v1/:assert/launchs/:id", controllers.DeleteLaunch)

	r.Run(":3000")
}
