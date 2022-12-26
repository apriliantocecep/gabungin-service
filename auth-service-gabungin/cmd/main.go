package main

import (
	"fmt"
	"log"
	"net"

	"github.com/apriliantocecep/auth-service-gabungin/pkg/config"
	"github.com/apriliantocecep/auth-service-gabungin/pkg/db"
	"github.com/apriliantocecep/auth-service-gabungin/pkg/pb"
	"github.com/apriliantocecep/auth-service-gabungin/pkg/services"
	"github.com/apriliantocecep/auth-service-gabungin/pkg/utils"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(&c)

	jwt := utils.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "auth-service-gabungin",
		ExpirationHours: 24 * 365,
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Auth Svc on", c.Port)

	s := services.Server{
		H:   h,
		Jwt: jwt,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
