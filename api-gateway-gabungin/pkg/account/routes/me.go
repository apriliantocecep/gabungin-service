package routes

import (
	"context"
	"net/http"

	"github.com/apriliantocecep/api-gateway-gabungin/pkg/account/pb"
	"github.com/gin-gonic/gin"
)

func Me(ctx *gin.Context, c pb.AccountServiceClient) {
	userId, _ := ctx.Get("userId")

	res, err := c.Read(context.Background(), &pb.ReadRequest{
		Id: userId.(int64),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
