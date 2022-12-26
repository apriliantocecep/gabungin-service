package account

import (
	"fmt"

	"github.com/apriliantocecep/api-gateway-gabungin/pkg/account/pb"
	"github.com/apriliantocecep/api-gateway-gabungin/pkg/account/routes"
	"github.com/apriliantocecep/api-gateway-gabungin/pkg/config"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.AccountServiceClient
	pb.UnimplementedAccountServiceServer
}

func InitServiceClient(c *config.Config) pb.AccountServiceClient {
	cc, err := grpc.Dial(c.AccountServiceurl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAccountServiceClient(cc)
}

func (svc *ServiceClient) Me(ctx *gin.Context) {
	routes.Me(ctx, svc.Client)
}

func (svc *ServiceClient) Create(ctx *gin.Context) {
	routes.Create(ctx, svc.Client)
}

func (svc *ServiceClient) Read(ctx *gin.Context) {
	routes.Read(ctx, svc.Client)
}

func (svc *ServiceClient) Update(ctx *gin.Context) {
	routes.Update(ctx, svc.Client)
}

func (svc *ServiceClient) Delete(ctx *gin.Context) {
	routes.Delete(ctx, svc.Client)
}

func (svc *ServiceClient) Browse(ctx *gin.Context) {
	routes.Browse(ctx, svc.Client)
}
