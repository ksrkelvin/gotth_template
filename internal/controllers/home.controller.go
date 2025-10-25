package controllers

import (
	"encontradev/views/pages"

	"github.com/gin-gonic/gin"
)

func HomeController(c *gin.Context) {
	homePage := pages.Home("Home")
	homePage.Render(c, c.Writer)
}
