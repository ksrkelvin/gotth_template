package controllers

import (
	"encontradev/internal/dto"
	"encontradev/views/pages"

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

	explorerPage := pages.Explorer(dto.UserResponse{})
	explorerPage.Render(ctx, ctx.Writer)
}
