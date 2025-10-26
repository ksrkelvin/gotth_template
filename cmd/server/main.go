package main

import (
	"encontradev/config"
	"encontradev/internal/controllers"
	"log"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	staticPath, _ := filepath.Abs("static")
	r.Static("/static", staticPath)

	diino, err := config.Init(r)
	if err != nil {
		log.Fatal("❌ Erro ao tentar inciar diino: ", err.Error())
	}

	controllers.RegisterControllers(r, diino)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal("❌ Erro ao iniciar encontraDev: ", err.Error())
	}

	log.Println("🚀 Servidor rodando em http://localhost:8080")

}
