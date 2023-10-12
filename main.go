package main

import (
	"context"
	"fmt"
	"log"
	"service_b/pb"

	"google.golang.org/grpc"
)

func call(s pb.ServiceAClient) {
	res, err := s.CreateUser(context.Background(), &pb.CreateUserRequest{
		Username: "nguyenhung",
		FullName: "hung",
		Email:    "hung@gmail.com",
		Password: "123456",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.GetUser())

}

func main() {
	clientConn, err := grpc.Dial("0.0.0.0:3003", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer clientConn.Close()

	client := pb.NewServiceAClient(clientConn)
	call(client)

	log.Printf("%f", client)

}
