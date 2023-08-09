package main

import (
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/pirosiki197/sodan-grpc/pkg/config"
	"github.com/pirosiki197/sodan-grpc/pkg/container"
	server "github.com/pirosiki197/sodan-grpc/pkg/grpc"
	"github.com/pirosiki197/sodan-grpc/pkg/grpc/pb"
	migration "github.com/pirosiki197/sodan-grpc/pkg/migraiton"
	"github.com/pirosiki197/sodan-grpc/pkg/repository"
	"github.com/pirosiki197/sodan-grpc/pkg/service"
	"google.golang.org/grpc"
)

func main() {
	listner, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	pb.RegisterAPIServiceServer(s, newServer())

	go func() {
		log.Println("Starting server on port :8080")
		s.Serve(listner)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Stopping server...")
	s.GracefulStop()
}

func newServer() pb.APIServiceServer {
	config := config.NewConfig()
	repository := repository.NewRepository(config)
	container := container.NewContainer(repository, config)
	migration.CreateDB(container)
	ss := service.NewSodanService(container)
	rs := service.NewReplyService(container)
	server := server.NewServer(ss, rs)

	return server
}
