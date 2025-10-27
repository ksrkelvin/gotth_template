package controllers

import (
	"encontradev/views/pages"
	"strings"

	"github.com/gin-gonic/gin"
)

func (c *Controllers) HomeController() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	c.eng.GET("/", c.GetHomePage)

	return
}

func (c *Controllers) GetHomePage(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			err := r.(error)
			ctx.String(500, "Erro inesperado: "+err.Error())
		}
	}()

	user, err := c.service.GetUser(ctx)
	if err != nil {
		ctx.String(500, "Erro ao tentar obter user: "+err.Error())
		return
	}

	partial := strings.ToLower(ctx.GetHeader("HX-Request")) == "true"

	homePage := pages.Home(user, partial)
	err = homePage.Render(ctx, ctx.Writer)
	if err != nil {
		ctx.String(500, "Erro ao tentar renderizar pagina: "+err.Error())
	}
}
