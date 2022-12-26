package main

import (
	"fmt"
	"log"
	"net"

	"github.com/apriliantocecep/account-service-gabungin/pkg/config"
	"github.com/apriliantocecep/account-service-gabungin/pkg/db"
	"github.com/apriliantocecep/account-service-gabungin/pkg/pb"
	"github.com/apriliantocecep/account-service-gabungin/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(&c)

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Family service on", c.Port)

	s := services.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAccountServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
