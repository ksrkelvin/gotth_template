package controllers

import (
	"encontradev/internal/dto"
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
		me.PUT("/")
		me.PUT("/avatar", c.UpdateAvatar)
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

func (c *Controllers) UpdateUser(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			err := r.(error)
			ctx.String(500, "Erro inesperado: "+err.Error())
		}
	}()

	var req = dto.UserUpdateRequest{}

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.String(400, "Body inválido")
		return
	}
	user, err := c.service.UpdateUser(ctx, req)
	if err != nil {
		ctx.String(500, "Falha ao atualizar usuário")
		return
	}
	ctx.Set("user", user)

	partial := ctx.GetHeader("HX-Request") == "true"

	mePage := pages.Me(user, partial)
	err = mePage.Render(ctx, ctx.Writer)
	if err != nil {
		ctx.String(500, "Erro ao tentar renderizar pagina: "+err.Error())
	}
}

func (c *Controllers) UpdateAvatar(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			err := r.(error)
			ctx.String(500, "Erro inesperado: "+err.Error())
		}
	}()

	var req = dto.UserUpdateRequest{}

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.String(400, "Body inválido")
		return
	}
	user, err := c.service.UpdateUser(ctx, req)
	if err != nil {
		ctx.String(500, "Falha ao atualizar usuário")
		return
	}
	ctx.Set("user", user)
	ctx.Header("HX-Trigger", "userUpdated")
	ctx.Status(200)
}
