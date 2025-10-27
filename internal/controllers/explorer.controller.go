package controllers

import (
	"encontradev/views/pages"
	"strings"

	"github.com/gin-gonic/gin"
)

func (c *Controllers) ExplorerController() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	explorer := c.eng.Group("/explorer")
	{
		explorer.GET("/", c.GetExplorerPage)
	}
	return
}

func (c *Controllers) GetExplorerPage(ctx *gin.Context) {
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

	partial := strings.ToLower(ctx.GetHeader("HX-Request")) == "true"

	explorerPage := pages.Explorer(user, partial)
	err = explorerPage.Render(ctx, ctx.Writer)
	if err != nil {
		ctx.String(500, "Erro ao tentar renderizar pagina: "+err.Error())
	}
}
