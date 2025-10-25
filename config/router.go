package config

import (
	"encontradev/internal/controllers"

	"github.com/gin-gonic/gin"
)

func (diino *Diino) RegisterRoutes(r *gin.Engine) {
	r.GET("/", controllers.HomeController)
	r.GET("/explorer", controllers.ExplorerController)
	r.GET("/notifications", controllers.NotificationsController)

}
