package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/apriliantocecep/api-gateway-gabungin/pkg/account/pb"
	"github.com/gin-gonic/gin"
)

func Browse(ctx *gin.Context, c pb.AccountServiceClient) {
	page, ok := ctx.GetQuery("page")
	if !ok {
		page = "1"
	}
	pageNumber, _ := strconv.ParseInt(page, 10, 32)

	pageSize, ok := ctx.GetQuery("page_size")
	if !ok {
		pageSize = "10"
	}
	pageSizeNumber, _ := strconv.ParseInt(pageSize, 10, 32)

	res, err := c.Browse(context.Background(), &pb.BrowseRequest{
		Page:     int32(pageNumber),
		PageSize: int32(pageSizeNumber),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
