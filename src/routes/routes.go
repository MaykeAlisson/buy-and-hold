package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/maykealisson/buy-and-hold/src/controllers"
	"github.com/maykealisson/buy-and-hold/src/middlewares"
)

func HandlerRequest() {
	r := gin.Default()

	r.GET("/", controllers.Swagger)

	// Auth
	r.POST("/api/v1/auth", controllers.Auth)

	// User
	r.POST("/api/v1/users", controllers.CreateUser)
	r.PUT("/api/v1/users/:id", middlewares.Auth(), controllers.UpdateUser)
	r.DELETE("/api/v1/users/:id", middlewares.Auth(), controllers.DeleteUser)

	// Assert
	r.GET("/api/v1/asserts", middlewares.Auth(), controllers.GetAssertBy)
	r.POST("/api/v1/asserts", middlewares.Auth(), controllers.CreateAssert)
	r.PUT("/api/v1/asserts/:id", middlewares.Auth(), controllers.UpdateAssert)
	r.DELETE("/api/v1/asserts/:id", middlewares.Auth(), controllers.DeleteAssert)

	// Launch
	r.GET("/api/v1/launchs/:month", middlewares.Auth(), controllers.GetByMonth)
	r.GET("/api/v1/:assert/launchs", middlewares.Auth(), controllers.GetByAssert)
	r.POST("/api/v1/:assert/launchs", middlewares.Auth(), controllers.CreateLaunch)
	r.PUT("/api/v1/:assert/launchs/:id", middlewares.Auth(), controllers.UpdateLaunch)
	r.DELETE("/api/v1/:assert/launchs/:id", middlewares.Auth(), controllers.DeleteLaunch)

	r.Run(":3000")
}
