package services

import (
	"context"
	"net/http"

	"github.com/apriliantocecep/account-service-gabungin/pkg/db"
	"github.com/apriliantocecep/account-service-gabungin/pkg/models"
	"github.com/apriliantocecep/account-service-gabungin/pkg/pb"
	"github.com/apriliantocecep/account-service-gabungin/utils"
)

type Server struct {
	H db.Handler
	pb.UnimplementedAccountServiceServer
}

func (s *Server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	var user models.User

	if result := s.H.DB.Where(&models.User{Email: in.Data.Email}).First(&user); result.Error == nil {
		return &pb.CreateResponse{
			Status: http.StatusConflict,
			Error:  "E-Mail already exists",
		}, nil
	}

	user.Email = in.Data.Email
	user.FirstName = in.Data.FirstName
	user.LastName = in.Data.LastName
	user.Gender = in.Data.Gender
	user.Status = int(in.Data.Status)
	user.Username = in.Data.Username
	user.Password = utils.HashPassword(in.Data.Password)

	if result := s.H.DB.Create(&user); result.Error != nil {
		return &pb.CreateResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	data := pb.User{
		Id:        user.Id,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Gender:    user.Gender,
		Status:    int32(user.Status),
		Username:  user.Username,
	}

	return &pb.CreateResponse{
		Status: http.StatusOK,
		Data:   &data,
	}, nil
}

func (s *Server) Read(ctx context.Context, in *pb.ReadRequest) (*pb.ReadResponse, error) {
	var user models.User

	if result := s.H.DB.First(&user, in.Id); result.Error != nil {
		return &pb.ReadResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	data := pb.User{
		Id:        user.Id,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Gender:    user.Gender,
		Status:    int32(user.Status),
		Username:  user.Username,
	}

	return &pb.ReadResponse{
		Status: http.StatusOK,
		Data:   &data,
	}, nil
}

func (s *Server) Update(ctx context.Context, in *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	var user models.User

	if result := s.H.DB.First(&user, in.Id); result.Error != nil {
		return &pb.UpdateResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	if result := s.H.DB.Where(&models.User{Email: in.Data.Email}).First(&user); result.Error == nil {
		return &pb.UpdateResponse{
			Status: http.StatusConflict,
			Error:  "E-Mail already exists",
		}, nil
	}

	if result := s.H.DB.Model(&user).Updates(models.User{
		FirstName: in.Data.FirstName,
		LastName:  in.Data.LastName,
		Gender:    in.Data.Gender,
		Username:  in.Data.Username,
		Status:    int(in.Data.Status),
		Password:  utils.HashPassword(in.Data.Password),
		Email:     in.Data.Email,
	}); result.Error != nil {
		return &pb.UpdateResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	data := pb.User{
		Id:        user.Id,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Gender:    user.Gender,
		Status:    int32(user.Status),
		Username:  user.Username,
	}

	return &pb.UpdateResponse{
		Status: http.StatusOK,
		Data:   &data,
	}, nil
}

func (s *Server) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	var user models.User

	if result := s.H.DB.First(&user, in.Id); result.Error != nil {
		return &pb.DeleteResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	s.H.DB.Delete(&models.User{}, user.Id)

	return &pb.DeleteResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *Server) Browse(ctx context.Context, in *pb.BrowseRequest) (*pb.BrowseResponse, error) {
	var users []models.User
	var datas []*pb.User

	page := in.Page
	pageSize := in.PageSize

	if result := s.H.DB.Scopes(utils.Paginate(int(page), int(pageSize))).Find(&users); result.Error != nil {
		return &pb.BrowseResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	for _, user := range users {
		data := pb.User{
			Id:        user.Id,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Gender:    user.Gender,
			Status:    int32(user.Status),
			Username:  user.Username,
		}
		datas = append(datas, &data)
	}

	return &pb.BrowseResponse{
		Status:   http.StatusOK,
		Data:     datas,
		Page:     page,
		PageSize: pageSize,
	}, nil
}
