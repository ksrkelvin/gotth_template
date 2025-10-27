package controllers

import (
	"encontradev/views/pages"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (c *Controllers) MeController() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	me := c.eng.Group("/me")
	{
		me.GET("/", c.GetMePage)
		me.PUT("/", c.UpdateMe)
	}

	return
}

func (c *Controllers) GetMePage(ctx *gin.Context) {
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

	fmt.Println(user)

	partial := ctx.GetHeader("HX-Request") == "true"

	mePage := pages.Me(user, partial)
	err = mePage.Render(ctx, ctx.Writer)
	if err != nil {
		ctx.String(500, "Erro ao tentar renderizar pagina: "+err.Error())
	}

}
func (c *Controllers) UpdateMe(ctx *gin.Context) {

}
