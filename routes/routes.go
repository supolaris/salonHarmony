package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/supolaris/salonHarmony/controller"
)

func InitRoutes(server *gin.Engine) {
	server.POST("createUser", controller.CreateUser)
	server.POST("createService", controller.CreateService)
}
