package controllers

import (
	"encontradev/views/pages"

	"github.com/gin-gonic/gin"
)

func NotificationsController(c *gin.Context) {
	partial := c.GetHeader("HX-Request") == "true"
	notificationsPage := pages.Notifications(partial)
	notificationsPage.Render(c, c.Writer)
}
