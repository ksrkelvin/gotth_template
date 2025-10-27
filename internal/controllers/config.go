package controllers

import (
	"encontradev/internal/auth"

	"github.com/gin-gonic/gin"
)

type Controllers struct {
	auth *auth.Auth
	eng  *gin.Engine
}

func RegisterControllers(eng *gin.Engine, auth *auth.Auth) (controllers *Controllers, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	c := &Controllers{
		eng:  eng,
		auth: auth,
	}

	c.AuthController()

	c.HomeController()
	c.ExplorerController()
	c.NotificationsController()

	return c, err
}
