package main

import (
	"context"
	pb "grpc_ty/gen/proto"
	"net"

	"log"

	"google.golang.org/grpc"
)

type testApiServer struct {
	pb.UnimplementedTestApiServer
}

func (s *testApiServer) Echo(ctx context.Context, req *pb.ResponseRequest) (*pb.ResponseRequest, error) {
	return req, nil
}

func (s *testApiServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{}, nil
}
func main() {

	log.Println("The gRPC server started...")
	// Create a tcp server
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Printf("Error with tcp connection")
	}

	// Create a gRPC server
	grpcServer := grpc.NewServer()

	// Register Server
	pb.RegisterTestApiServer(grpcServer, &testApiServer{})

	// Connect tcp and grpc
	if err := grpcServer.Serve(listen); err != nil {
		log.Printf("cannot connect grpc server to tcp connection")
	}
}
