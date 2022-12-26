package db

import (
	"fmt"
	"log"

	"github.com/apriliantocecep/account-service-gabungin/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(c *config.Config) Handler {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", c.Db_host, c.Db_port, c.Db_user, c.Db_name, c.Db_password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	// db.AutoMigrate(&models.User{})

	return Handler{db}
}
