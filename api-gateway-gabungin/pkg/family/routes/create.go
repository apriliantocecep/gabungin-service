package routes

import (
	"context"
	"net/http"

	"github.com/apriliantocecep/api-gateway-gabungin/pkg/family/pb"
	"github.com/gin-gonic/gin"
)

type CreateFamilyRequest struct {
	ParentId uint32 `json:"parent_id"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Order    int32  `json:"order"`
}

func Create(ctx *gin.Context, c pb.FamilyServiceClient) {
	body := CreateFamilyRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId, _ := ctx.Get("userId")

	familyData := pb.DataFamily{
		UserId:   userId.(int64),
		ParentId: body.ParentId,
		Name:     body.Name,
		Gender:   body.Gender,
		Order:    body.Order,
	}

	res, err := c.Create(context.Background(), &pb.CreateRequest{
		Data: &familyData,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
