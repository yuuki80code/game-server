package router

import (
	"github/com/yuuki80code/game-server/controller"
	"github/com/yuuki80code/game-server/middleware"
)

var UserController = &controller.UserController{}

func initUserRouter() {
	apiRouter.POST("/register", UserController.Register)
	apiRouter.POST("/login", UserController.Login)
	apiRouter.GET("/user/info", middleware.CheckToken(), UserController.Userinfo)
}
