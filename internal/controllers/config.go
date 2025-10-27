package controllers

import (
	"encontradev/internal/auth"
	"encontradev/internal/service"

	"github.com/gin-gonic/gin"
)

type Controllers struct {
	auth    *auth.Auth
	service *service.Service
	eng     *gin.Engine
}

func RegisterControllers(eng *gin.Engine, auth *auth.Auth, service *service.Service) (controllers *Controllers, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	c := &Controllers{
		eng:     eng,
		auth:    auth,
		service: service,
	}

	eng.Use(auth.JWTMiddleware())

	err = c.AuthController()
	if err != nil {
		return
	}

	err = c.HomeController()
	if err != nil {
		return
	}
	err = c.ExplorerController()
	if err != nil {
		return
	}
	err = c.NotificationsController()
	if err != nil {
		return
	}
	err = c.MeController()
	if err != nil {
		return
	}

	return c, err
}
