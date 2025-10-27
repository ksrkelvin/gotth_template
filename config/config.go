package config

import (
	"encontradev/internal/auth"
	"encontradev/internal/controllers"
	"encontradev/internal/repository"
	"encontradev/internal/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Diino struct {
	DB          *gorm.DB
	Routes      *gin.Engine
	Auth        *auth.Auth
	Controllers *controllers.Controllers
	Repository  *repository.Repository
	Service     *service.Service
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

	d.Repository, err = repository.RegisterRepository(d.DB)
	if err != nil {
		log.Fatal("⚠️  Não foi possivel inicializar sistemas de repository")
	}

	d.Auth, err = auth.SetAuth(d.Repository)
	if err != nil {
		log.Fatal("⚠️  Não foi possivel inicializar sistemas de auth")
	}

	d.Service, err = service.RegisterService(d.Repository)
	if err != nil {
		log.Fatal("⚠️  Não foi possivel inicializar sistemas de services")
	}

	return d, err
}
