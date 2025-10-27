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

	diino.Controllers, err = controllers.RegisterControllers(r, diino.Auth)
	if err != nil {
		log.Fatal("⚠️  Não foi possivel inicializar sistemas de services")
	}

	err = r.Run(":8080")
	if err != nil {
		log.Fatal("❌ Erro ao iniciar server: ", err.Error())
	}

	log.Println("🚀 Servidor rodando em http://localhost:8080")

}
