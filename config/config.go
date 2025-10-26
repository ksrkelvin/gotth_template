package config

import (
	"encontradev/internal/auth"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Diino struct {
	DB     *gorm.DB
	Routes *gin.Engine
	Auth   *auth.Auth
}

func Init(r *gin.Engine) (diino *Diino, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()
	if err := godotenv.Load(); err != nil {
		log.Fatal("⚠️  Nenhum .env encontrado, usando variáveis do ambiente")
	}

	d := &Diino{}
	d.ConnectDB()

	d.Auth, err = auth.SetAuth(d.DB)
	if err != nil {
		log.Fatal("⚠️  Não foi possivel inicializar sistemas de auth")
	}

	return d, err
}
