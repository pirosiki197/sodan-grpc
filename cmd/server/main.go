package main

import (
	"fmt"
	"net/http"

	"github.com/pirosiki197/sodan-grpc/pkg/config"
	"github.com/pirosiki197/sodan-grpc/pkg/container"
	server "github.com/pirosiki197/sodan-grpc/pkg/grpc"
	"github.com/pirosiki197/sodan-grpc/pkg/grpc/pb/api/v1/apiv1connect"
	migration "github.com/pirosiki197/sodan-grpc/pkg/migraiton"
	"github.com/pirosiki197/sodan-grpc/pkg/repository"
	"github.com/pirosiki197/sodan-grpc/pkg/service"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	fmt.Println("Starting server on port :8080")

	mux := http.NewServeMux()

	path, handler := apiv1connect.NewAPIServiceHandler(newServer())
	mux.Handle(path, handler)
	corsHandler := cors.AllowAll().Handler(h2c.NewHandler(mux, &http2.Server{}))

	fmt.Println("Server is ready to handle requests at", "http://localhost:8080")
	http.ListenAndServe(
		"localhost:8080",
		corsHandler,
	)
}

func newServer() apiv1connect.APIServiceHandler {
	config := config.NewConfig()
	repository := repository.NewRepository(config)
	container := container.NewContainer(repository, config)
	migration.CreateDB(container)
	ss := service.NewSodanService(container)
	rs := service.NewReplyService(container)
	server := server.NewServer(ss, rs)

	return server
}
