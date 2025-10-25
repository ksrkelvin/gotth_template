package controllers

import (
	"encontradev/views/pages"

	"github.com/gin-gonic/gin"
)

func ExplorerController(c *gin.Context) {
	partial := c.GetHeader("HX-Request") == "true"
	explorerPage := pages.Explorer(partial)
	explorerPage.Render(c, c.Writer)
}
