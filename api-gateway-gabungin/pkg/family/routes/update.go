package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/apriliantocecep/api-gateway-gabungin/pkg/family/pb"
	"github.com/gin-gonic/gin"
)

type UpdateFamilyRequest struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Order  int32  `json:"order"`
}

func Update(ctx *gin.Context, c pb.FamilyServiceClient) {
	body := UpdateFamilyRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	param_id := ctx.Param("id")

	id, _ := strconv.ParseInt(param_id, 10, 64)

	familyData := pb.DataFamily{
		Name:   body.Name,
		Gender: body.Gender,
		Order:  body.Order,
	}

	res, err := c.Update(context.Background(), &pb.UpdateRequest{
		Id:   id,
		Data: &familyData,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
