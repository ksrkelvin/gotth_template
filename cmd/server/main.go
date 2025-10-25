package main

import (
	"encontradev/config"
	"log"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	diino, err := config.Init()
	if err != nil {
		log.Fatal("❌ Erro ao tentar inciar diino: ", err.Error())
	}

	r := gin.Default()

	staticPath, _ := filepath.Abs("static")
	r.Static("/static", staticPath)

	diino.ConnectDB()
	diino.RegisterRoutes(r)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal("❌ Erro ao iniciar encontraDev: ", err.Error())
	}

	log.Println("🚀 Servidor rodando em http://localhost:8080")

}
