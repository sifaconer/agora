package services

import (
	pb "agora/api/proto/libgo"
	"agora/core/models"
	"agora/repository"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServicesFormat struct {
	Usecase repository.FormatRepository
	pb.StringFormatServiceServer
}

func (s *ServicesFormat) Reverse(ctx context.Context, in *pb.DataRequest) (*pb.DataResponse, error) {
	resp, err := s.Usecase.Reverse(&models.Transport{
		Value: in.Value,
	})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	return &pb.DataResponse{
		Value:     resp.Value,
		FuncApply: resp.FuncApply,
	}, nil
}

func (s *ServicesFormat) Upper(ctx context.Context, in *pb.DataRequest) (*pb.DataResponse, error) {
	resp, err := s.Usecase.Upper(&models.Transport{
		Value: in.Value,
	})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	return &pb.DataResponse{
		Value:     resp.Value,
		FuncApply: resp.FuncApply,
	}, nil
}

func (s *ServicesFormat) Lower(ctx context.Context, in *pb.DataRequest) (*pb.DataResponse, error) {
	resp, err := s.Usecase.Lower(&models.Transport{
		Value: in.Value,
	})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	return &pb.DataResponse{
		Value:     resp.Value,
		FuncApply: resp.FuncApply,
	}, nil
}
