package controllers

import (
	"encontradev/internal/dto"
	"encontradev/views/pages"

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

	homePage := pages.Home(dto.UserResponse{})
	homePage.Render(ctx, ctx.Writer)
}
