package main

import (
	"log"
	"net"

	"github.com/FellipeG/fullcycle-grpc-go/pb"
	"github.com/FellipeG/fullcycle-grpc-go/services"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
}