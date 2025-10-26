package controllers

import (
	"encontradev/config"

	"github.com/gin-gonic/gin"
)

func RegisterControllers(r *gin.Engine, diino *config.Diino) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	AuthController(r, diino)

	HomeController(r, diino)
	ExplorerController(r, diino)
	NotificationsController(r, diino)

	return
}
