package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/apriliantocecep/api-gateway-gabungin/pkg/account/pb"
	"github.com/gin-gonic/gin"
)

func Delete(ctx *gin.Context, c pb.AccountServiceClient) {
	param_id := ctx.Param("id")

	id, _ := strconv.ParseInt(param_id, 10, 64)

	res, err := c.Delete(context.Background(), &pb.DeleteRequest{
		Id: id,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
