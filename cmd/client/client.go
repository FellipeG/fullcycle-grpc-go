package main

import (
	"context"
	"log"

	"github.com/FellipeG/fullcycle-grpc-go/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	AddUser(client)


}

func AddUser(client pb.UserServiceClient) {

	req := &pb.User{
		Id: "0",
		Name: "Fellipe",
		Email: "fellipeg.rjqoor@gmail.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	log.Println(res)

}