package config

import (
	"encontradev/internal/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (diino *Diino) ConnectDB() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	pgUser := os.Getenv("POSTGRES_USER")
	pgPass := os.Getenv("POSTGRES_PASSWORD")
	pgDb := os.Getenv("POSTGRES_DB")

	pgHost := os.Getenv("POSTGRES_HOST")
	if pgHost == "" {
		pgHost = "localhost"
	}
	pgPort := os.Getenv("POSTGRES_PORT")
	if pgPort == "" {
		pgPort = "5432"
	}

	dbConnUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", pgUser, pgPass, pgHost, pgPort, pgDb)
	database, err := gorm.Open(postgres.Open(dbConnUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Falha ao conectar no banco:", err)
	}

	err = database.AutoMigrate(models.User{})
	if err != nil {
		log.Fatal("❌ Falha ao migrar tabelas:", err)
	}

	diino.DB = database
	log.Println("✅ Banco conectado com sucesso!")

	return err
}
