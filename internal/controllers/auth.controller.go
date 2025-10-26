package controllers

import (
	"encontradev/config"
	"encontradev/views/pages"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	diino *config.Diino
}

func AuthController(r *gin.Engine, diino *config.Diino) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	a := &Auth{
		diino: diino,
	}

	login := r.Group("/login")
	{
		login.GET("/", a.GetLoginPage)
	}

	auth := r.Group("/auth")
	{
		google := auth.Group("/google")
		{
			google.GET("/", a.diino.Auth.GoogleLogin)
			google.GET("/callback", a.diino.Auth.GoogleCallback)
		}
	}
	return
}

func (a *Auth) GetLoginPage(c *gin.Context) {
	loginPage := pages.Login()
	loginPage.Render(c, c.Writer)
}
