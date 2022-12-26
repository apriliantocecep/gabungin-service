package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/apriliantocecep/api-gateway-gabungin/pkg/family/pb"
	"github.com/gin-gonic/gin"
)

type family struct {
	ID       uint32    `json:"id"`
	Name     string    `json:"name"`
	Gender   string    `json:"gender"`
	Order    string    `json:"order"`
	Families []*family `json:"data"`
}

func All(ctx *gin.Context, c pb.FamilyServiceClient) {
	// param_id := ctx.Param("id")
	// id, _ := strconv.ParseInt(param_id, 10, 32)

	userId, _ := ctx.Get("userId")

	fams := getDataFamily(c, userId.(int64), 0)

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   fams,
	})
}

func getDataFamily(c pb.FamilyServiceClient, userId int64, parentId uint32) []*family {
	res, err := c.GetDataFamily(context.Background(), &pb.GetDataFamilyRequest{
		UserId: userId,
		Id:     parentId,
	})
	if err != nil {
		panic(err)
	}

	var fams []*family

	if len(res.Data) > 0 {
		for i, data := range res.Data {
			var fam family
			fam.ID = data.Id
			fam.Name = data.Name
			fam.Gender = data.Gender
			fam.Order = ordinalOrder(i + 1)
			fam.Families = getDataFamily(c, data.UserId, data.Id)
			fams = append(fams, &fam)
		}
	}

	return fams
}

func ordinalOrder(num int) string {
	return fmt.Sprintf("Anak ke %d", num)
}
