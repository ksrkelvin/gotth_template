package controllers

import (
	"encontradev/config"
	"encontradev/internal/dto"
	"encontradev/views/pages"

	"github.com/gin-gonic/gin"
)

type Home struct {
	diino *config.Diino
}

func HomeController(r *gin.Engine, diino *config.Diino) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	h := &Home{
		diino: diino,
	}

	r.GET("/", h.GetHomePage)

	return
}

func (e *Home) GetHomePage(c *gin.Context) {
	homePage := pages.Home(dto.User{})
	homePage.Render(c, c.Writer)
}
