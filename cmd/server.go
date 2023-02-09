package cmd

import (
	pb "agora/api/proto/libgo"
	"agora/api/services"
	"agora/core/format"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run() error {

	// server definition
	lis, err := net.Listen("tcp", net.JoinHostPort("0.0.0.0", "9009"))
	if err != nil {
		return err
	}

	// grpc server
	grpcServer := grpc.NewServer()
	pb.RegisterStringFormatServiceServer(grpcServer, &services.ServicesFormat{
		Usecase: format.NewFormat(),
	})

	// grpc reflect
	reflection.Register(grpcServer)

	// server running
	log.Printf("Server started at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		return err
	}

	return nil
}
