package main

import (
	"context"
	"fmt"
	"log"

	"github.com/pirosiki197/sodan-grpc/pkg/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	client pb.APIServiceClient
)

func main() {
	fmt.Println("start grpc client")

	address := "localhost:8080"
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal(err)
	}

	client = pb.NewAPIServiceClient(conn)

	GetSodan()
}

func GetSodan() {
	req := &pb.GetSodanRequest{
		Id: 1,
	}
	res, err := client.GetSodan(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.GetSodan())
}
