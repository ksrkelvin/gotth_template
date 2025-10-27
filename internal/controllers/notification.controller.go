package controllers

import (
	"encontradev/internal/dto"
	"encontradev/views/pages"

	"github.com/gin-gonic/gin"
)

func (c *Controllers) NotificationsController() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	notifications := c.eng.Group("/notifications")
	{
		notifications.GET("/", c.GetNotificationsPage)
	}
	return
}

func (c *Controllers) GetNotificationsPage(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			err := r.(error)
			ctx.String(500, "Erro inesperado: "+err.Error())
		}
	}()

	notificationsPage := pages.Notifications(dto.UserResponse{})
	notificationsPage.Render(ctx, ctx.Writer)
}
