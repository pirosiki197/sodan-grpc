package grpc

import (
	"os"

	"github.com/pirosiki197/sodan-grpc/pkg/service"
	"golang.org/x/exp/slog"
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
	go checkNewReply()
	return s
}
