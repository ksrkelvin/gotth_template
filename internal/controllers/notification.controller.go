package controllers

import (
	"encontradev/config"
	"encontradev/internal/dto"
	"encontradev/views/pages"

	"github.com/gin-gonic/gin"
)

type Notifications struct {
	diino *config.Diino
}

func NotificationsController(r *gin.Engine, diino *config.Diino) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	n := &Notifications{
		diino: diino,
	}

	home := r.Group("/notifications")
	{
		home.GET("/", n.GetNotificationsPage)
	}
	return
}

func (e *Notifications) GetNotificationsPage(c *gin.Context) {
	notificationsPage := pages.Notifications(dto.User{})
	notificationsPage.Render(c, c.Writer)
}
