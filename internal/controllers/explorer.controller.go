package controllers

import (
	"encontradev/config"
	"encontradev/internal/dto"
	"encontradev/views/pages"

	"github.com/gin-gonic/gin"
)

type Explorer struct {
	diino *config.Diino
}

func ExplorerController(r *gin.Engine, diino *config.Diino) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	e := &Explorer{
		diino: diino,
	}

	explorer := r.Group("/explorer")
	{
		explorer.GET("/", e.GetExplorerPage)
	}
	return
}

func (e *Explorer) GetExplorerPage(c *gin.Context) {
	explorerPage := pages.Explorer(dto.User{})
	explorerPage.Render(c, c.Writer)
}
