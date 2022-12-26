package services

import (
	"context"
	"net/http"

	"github.com/apriliantocecep/family-service-gabungin/pkg/db"
	"github.com/apriliantocecep/family-service-gabungin/pkg/models"
	"github.com/apriliantocecep/family-service-gabungin/pkg/pb"
)

type Server struct {
	H db.Handler
	pb.UnimplementedFamilyServiceServer
}

func (s *Server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	var family models.Family

	family.UserId = in.Data.UserId
	family.ParentId = in.Data.ParentId
	family.Name = in.Data.Name
	family.Gender = in.Data.Gender
	family.Order = in.Data.Order

	if result := s.H.DB.Create(&family); result.Error != nil {
		return &pb.CreateResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.CreateResponse{
		Status: http.StatusOK,
		Id:     int64(family.ID),
	}, nil
}

func (s *Server) Read(ctx context.Context, in *pb.ReadRequest) (*pb.ReadResponse, error) {
	var family models.Family

	if result := s.H.DB.First(&family, in.Id); result.Error != nil {
		return &pb.ReadResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	familyData := pb.DataFamily{
		Id:       uint32(family.ID),
		UserId:   family.UserId,
		ParentId: family.ParentId,
		Name:     family.Name,
		Gender:   family.Gender,
		Order:    family.Order,
	}

	return &pb.ReadResponse{
		Status: http.StatusOK,
		Data:   &familyData,
	}, nil
}

func (s *Server) Update(ctx context.Context, in *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	var family models.Family

	if result := s.H.DB.First(&family, in.Id); result.Error != nil {
		return &pb.UpdateResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	if result := s.H.DB.Model(&family).Updates(models.Family{
		Name:   in.Data.Name,
		Gender: in.Data.Gender,
		Order:  in.Data.Order,
	}); result.Error != nil {
		return &pb.UpdateResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	familyData := pb.DataFamily{
		Id:       uint32(family.ID),
		UserId:   family.UserId,
		ParentId: family.ParentId,
		Name:     family.Name,
		Gender:   family.Gender,
		Order:    family.Order,
	}

	return &pb.UpdateResponse{
		Status: http.StatusOK,
		Data:   &familyData,
	}, nil
}

func (s *Server) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	var family models.Family

	if result := s.H.DB.First(&family, in.Id); result.Error != nil {
		return &pb.DeleteResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	s.H.DB.Delete(&models.Family{}, family.ID)

	return &pb.DeleteResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *Server) GetDataFamily(ctx context.Context, in *pb.GetDataFamilyRequest) (*pb.GetDataFamilyResponse, error) {
	var families []models.Family
	var familiesData []*pb.DataFamily

	if result := s.H.DB.Where(map[string]interface{}{"user_id": in.UserId, "parent_id": in.Id}).Find(&families); result.Error != nil {
		return &pb.GetDataFamilyResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	for _, family := range families {
		familyData := pb.DataFamily{
			Id:       uint32(family.ID),
			UserId:   family.UserId,
			ParentId: family.ParentId,
			Name:     family.Name,
			Gender:   family.Gender,
			Order:    family.Order,
		}
		familiesData = append(familiesData, &familyData)
	}

	return &pb.GetDataFamilyResponse{
		Status: http.StatusOK,
		Data:   familiesData,
	}, nil
}
