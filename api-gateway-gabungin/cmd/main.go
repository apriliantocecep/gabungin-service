package main

import (
	"log"

	"github.com/apriliantocecep/api-gateway-gabungin/pkg/account"
	"github.com/apriliantocecep/api-gateway-gabungin/pkg/auth"
	"github.com/apriliantocecep/api-gateway-gabungin/pkg/config"
	"github.com/apriliantocecep/api-gateway-gabungin/pkg/family"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()
	r.Use(cors.Default())

	authService := *auth.RegisterRoutes(r, &c)
	family.RegisterRoutes(r, &c, &authService)
	account.RegisterRoutes(r, &c, &authService)

	r.Run(c.Port)
}
