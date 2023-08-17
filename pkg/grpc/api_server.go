package grpc

import (
	"os"

	"log/slog"

	"github.com/pirosiki197/sodan-grpc/pkg/service"
)

type server struct {
	sodanService service.SodanService
	replyService service.ReplyService
	logger       *slog.Logger
}

func NewServer(ss service.SodanService, rs service.ReplyService) *server {
	s := &server{
		sodanService: ss,
		replyService: rs,
		logger:       slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true})),
	}
	go s.checkNewReply()
	return s
}
