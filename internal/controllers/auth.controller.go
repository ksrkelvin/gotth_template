package controllers

import (
	"encontradev/internal/dto"
	"encontradev/views/pages"

	"github.com/gin-gonic/gin"
)

func (c *Controllers) AuthController() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	c.eng.GET("/login", c.GetLoginPage)
	c.eng.POST("/logout", c.Logout)

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

func (a *Controllers) Logout(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			err := r.(error)
			ctx.String(500, "Erro inesperado: "+err.Error())
		}
	}()

	ctx.SetCookie(
		"jwt",
		"",
		-1,
		"/",
		"",
		false,
		true,
	)

	ctx.Header("HX-Redirect", "/")
	ctx.Status(200)
}

func (a *Controllers) GetLoginPage(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			err := r.(error)
			ctx.String(500, "Erro inesperado: "+err.Error())
		}
	}()

	partial := ctx.GetHeader("HX-Request") == "true"

	loginPage := pages.Login(dto.UserResponse{}, partial)
	err := loginPage.Render(ctx, ctx.Writer)
	if err != nil {
		ctx.String(500, "Erro ao tentar renderizar pagina: "+err.Error())
	}
}
