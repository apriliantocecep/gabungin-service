package account

import (
	"github.com/apriliantocecep/api-gateway-gabungin/pkg/auth"
	"github.com/apriliantocecep/api-gateway-gabungin/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine, c *config.Config, authService *auth.ServiceClient) *ServiceClient {
	a := auth.InitAuthMiddleware(authService)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	r := e.Group("/account").Use(a.AuthRequired)
	{
		r.GET("/browse", svc.Browse)
		r.GET("/me", svc.Me)
		r.POST("/create", svc.Create)
		r.GET("/read/:id", svc.Read)
		r.PUT("/update/:id", svc.Update)
		r.DELETE("/delete/:id", svc.Delete)
	}

	return svc
}
