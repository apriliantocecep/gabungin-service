package family

import (
	"fmt"

	"github.com/apriliantocecep/api-gateway-gabungin/pkg/config"
	"github.com/apriliantocecep/api-gateway-gabungin/pkg/family/pb"
	"github.com/apriliantocecep/api-gateway-gabungin/pkg/family/routes"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.FamilyServiceClient
}

func InitServiceClient(c *config.Config) pb.FamilyServiceClient {
	cc, err := grpc.Dial(c.FamilyServiceurl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewFamilyServiceClient(cc)
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

func (svc *ServiceClient) All(ctx *gin.Context) {
	routes.All(ctx, svc.Client)
}
