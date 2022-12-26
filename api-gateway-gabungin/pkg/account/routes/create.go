package routes

import (
	"context"
	"net/http"

	"github.com/apriliantocecep/api-gateway-gabungin/pkg/account/pb"
	"github.com/gin-gonic/gin"
)

type CreateAccountRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Status    int    `json:"status"`
}

func Create(ctx *gin.Context, c pb.AccountServiceClient) {
	body := CreateAccountRequest{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// userId, _ := ctx.Get("userId")

	data := pb.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Gender:    body.Gender,
		Status:    int32(body.Status),
		Username:  body.Username,
		Email:     body.Email,
		Password:  body.Password,
	}

	res, err := c.Create(context.Background(), &pb.CreateRequest{
		Data: &data,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
