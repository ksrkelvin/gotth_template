package config

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Diino struct {
	DB     *gorm.DB
	Routes *gin.Engine
}

func Init() (diino *Diino, err error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("⚠️  Nenhum .env encontrado, usando variáveis do ambiente")
	}

	return &Diino{}, err
}
