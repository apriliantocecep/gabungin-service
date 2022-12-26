package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/apriliantocecep/api-gateway-gabungin/pkg/account/pb"
	"github.com/gin-gonic/gin"
)

func Read(ctx *gin.Context, c pb.AccountServiceClient) {
	param_id := ctx.Param("id")

	id, _ := strconv.ParseInt(param_id, 10, 64)

	res, err := c.Read(context.Background(), &pb.ReadRequest{
		Id: id,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
