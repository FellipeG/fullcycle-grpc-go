package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	//AddUserVerbose(client)
	//AddUser(client)
	//AddUsers(client)
	AddUserStreamBoth(client)


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

func AddUserVerbose(client pb.UserServiceClient) {

	req := &pb.User{
		Id: "0",
		Name: "Fellipe",
		Email: "fellipeg.rjqoor@gmail.com",
	}

	resStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := resStream.Recv()
		if err == io.EOF {
			break;
		}

		if err != nil {
			log.Fatalf("Could not receive the message: %v", err)
		}

		fmt.Println("Status: ", stream.Status , " - ", stream.GetUser())
	}

}

func AddUsers(client pb.UserServiceClient) {

	reqs := []*pb.User{
		&pb.User{
			Id: "f1",
			Name: "Fellipe",
			Email: "fel1@email.com",
		},
		&pb.User{
			Id: "f2",
			Name: "Fellipe 2",
			Email: "fel2@email.com",
		},
		&pb.User{
			Id: "f3",
			Name: "Fellipe 3",
			Email: "fel3@email.com",
		},
		&pb.User{
			Id: "f4",
			Name: "Fellipe 4",
			Email: "fel4@email.com",
		},
		&pb.User{
			Id: "f5",
			Name: "Fellipe 5",
			Email: "fel5@email.com",
		},
		&pb.User{
			Id: "f6",
			Name: "Fellipe 6",
			Email: "fel6@email.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)

}

func AddUserStreamBoth(client pb.UserServiceClient) {

	stream, err := client.AddUserStreamBoth(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	reqs := []*pb.User{
		&pb.User{
			Id: "f1",
			Name: "Fellipe",
			Email: "fel1@email.com",
		},
		&pb.User{
			Id: "f2",
			Name: "Fellipe 2",
			Email: "fel2@email.com",
		},
		&pb.User{
			Id: "f3",
			Name: "Fellipe 3",
			Email: "fel3@email.com",
		},
		&pb.User{
			Id: "f4",
			Name: "Fellipe 4",
			Email: "fel4@email.com",
		},
		&pb.User{
			Id: "f5",
			Name: "Fellipe 5",
			Email: "fel5@email.com",
		},
		&pb.User{
			Id: "f6",
			Name: "Fellipe 6",
			Email: "fel6@email.com",
		},
	}

	wait := make(chan int)

	go func() {

		for _, req := range reqs {
			fmt.Println("Sending user: ", req.GetName())
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}

		stream.CloseSend()

	}()

	go func() {
		
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error receiving response: %v", err)
				break
			}

			fmt.Printf("Recebendo user %v, com status: %v\n", res.GetUser().GetName(), res.GetStatus())
		}

		close(wait)

	}()

	<-wait
}