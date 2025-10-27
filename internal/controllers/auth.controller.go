package controllers

import (
	"encontradev/views/pages"

	"github.com/gin-gonic/gin"
)

func (c *Controllers) AuthController() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	login := c.eng.Group("/login")
	{
		login.GET("/", c.GetLoginPage)
	}

	auth := c.eng.Group("/auth")
	{
		google := auth.Group("/google")
		{
			google.GET("/", c.auth.GoogleLogin)
			google.GET("/callback", c.auth.GoogleCallback)
		}
	}
	return
}

func (a *Controllers) GetLoginPage(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			err := r.(error)
			ctx.String(500, "Erro inesperado: "+err.Error())
		}
	}()

	loginPage := pages.Login()
	loginPage.Render(ctx, ctx.Writer)
}
