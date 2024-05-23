package routers

import (
	"clicks/controllers"
	"clicks/middlewares"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	users := router.Group("/users")

	users.POST("/register", controllers.RegisterUserController)
	users.POST("/login", controllers.LoginUserController)
	users.POST("/profiles", middlewares.Authentication(), controllers.ViewProfiles)
	users.POST("/swipe", middlewares.Authentication(), controllers.ViewProfiles)



	return router
}
