package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"connectrpc.com/connect"
	apiv1 "github.com/pirosiki197/sodan-grpc/pkg/grpc/pb/api/v1"
	"github.com/pirosiki197/sodan-grpc/pkg/grpc/pb/api/v1/apiv1connect"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	scanner *bufio.Scanner
)

func main() {
	fmt.Println("start grpc client")

	scanner = bufio.NewScanner(os.Stdin)

	client := apiv1connect.NewAPIServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	)

	for {
		fmt.Println("1: GetSodan")
		fmt.Println("2: GetSodanList")
		fmt.Println("3: CreateReply")
		fmt.Println("4: SubscribeSodan")
		fmt.Println("5: Exit")

		scanner.Scan()
		in := scanner.Text()

		switch in {
		case "1":
			GetSodan(client)
		case "2":
			GetSodanList(client)
		case "3":
			CreateReply(client)
		case "4":
			SubscribeSodan(client)
		case "5":
			fmt.Println("exit")
			goto M
		}
	}
M:
}

func GetSodan(client apiv1connect.APIServiceClient) {
	fmt.Println("GetSodan")
	req := connect.NewRequest(&apiv1.GetSodanRequest{
		Id: 1,
	})
	res, err := client.GetSodan(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Msg.Sodan)
}

func GetSodanList(client apiv1connect.APIServiceClient) {
	fmt.Println("GetSodanList")
	req := connect.NewRequest(&emptypb.Empty{})
	res, err := client.GetSodanList(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Msg.Sodans)
}

func CreateReply(client apiv1connect.APIServiceClient) {
	fmt.Println("CreateReply")
	req := connect.NewRequest(&apiv1.CreateReplyRequest{
		Text:      "reply",
		CreaterId: "pirosiki",
		SodanId:   1,
	})
	res, err := client.CreateReply(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Msg.Id)
}

func SubscribeSodan(client apiv1connect.APIServiceClient) {
	fmt.Println("SubscribeSodan")
	req := connect.NewRequest(&apiv1.SubscribeSodanRequest{
		Id: 1,
	})
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	stream, err := client.SubscribeSodan(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		i := 0
		for stream.Receive() {
			if i == 3 {
				cancel()
				return
			}
			res := stream.Msg()
			fmt.Println(res.Reply)
			i++
		}
	}()
}
