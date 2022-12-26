package auth

import (
	"github.com/apriliantocecep/api-gateway-gabungin/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	r := e.Group("/auth")
	{
		r.POST("/register", svc.Register)
		r.POST("/login", svc.Login)
	}

	return svc
}
