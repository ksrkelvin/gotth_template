package controllers

import (
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

	user, err := c.service.GetUser(ctx)
	if err != nil {
		ctx.String(500, "Erro ao tentar obter user: "+err.Error())
	}

	partial := ctx.GetHeader("HX-Request") == "true"

	notificationsPage := pages.Notifications(user, partial)
	err = notificationsPage.Render(ctx, ctx.Writer)
	if err != nil {
		ctx.String(500, "Erro ao tentar renderizar pagina: "+err.Error())
	}
}
