package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/apriliantocecep/api-gateway-gabungin/pkg/account/pb"
	"github.com/gin-gonic/gin"
)

type UpdateAccountRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Status    int    `json:"status"`
}

func Update(ctx *gin.Context, c pb.AccountServiceClient) {
	body := UpdateAccountRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	param_id := ctx.Param("id")

	id, _ := strconv.ParseInt(param_id, 10, 64)

	data := pb.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Gender:    body.Gender,
		Status:    int32(body.Status),
		Username:  body.Username,
		Email:     body.Email,
		Password:  body.Password,
	}

	res, err := c.Update(context.Background(), &pb.UpdateRequest{
		Id:   id,
		Data: &data,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
