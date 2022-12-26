package db

import (
	"fmt"
	"log"

	"github.com/apriliantocecep/auth-service-gabungin/pkg/config"
	"github.com/apriliantocecep/auth-service-gabungin/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(c *config.Config) Handler {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", c.Db_host, c.Db_port, c.Db_user, c.Db_name, c.Db_password)
	db, err := gorm.Open(postgres.Open(DBURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.User{})

	return Handler{db}
}
