package grpc

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pirosiki197/sodan-grpc/pkg/grpc/pb"
	"github.com/pirosiki197/sodan-grpc/pkg/repository/model/dto"
	"github.com/pirosiki197/sodan-grpc/pkg/service"
)

type server struct {
	pb.UnimplementedAPIServiceServer
	sodanService service.SodanService
	replyService service.ReplyService
}

func NewServer(ss service.SodanService, rs service.ReplyService) *server {
	return &server{
		sodanService: ss,
		replyService: rs,
	}
}

func (s *server) CreateSodan(ctx context.Context, req *pb.CreateSodanRequest) (*pb.CreateSodanResponse, error) {
	sodan := &dto.SodanDto{
		Title:     req.GetTitle(),
		Text:      req.GetText(),
		CreaterID: req.GetCreaterID(),
	}

	id, err := s.sodanService.CreateSodan(sodan)
	if err != nil {
		return nil, err
	}

	res := &pb.CreateSodanResponse{
		Id: uint64(id),
	}
	return res, nil
}

func (s *server) GetSodan(ctx context.Context, req *pb.GetSodanRequest) (*pb.GetSodanResponse, error) {
	id := req.GetId()
	sodan, err := s.sodanService.FindByID(uint(id))
	if err != nil {
		return nil, err
	}
	// TODO: sodanをpb.Sodanに変換する関数を作成する -> pbconv
	res := &pb.GetSodanResponse{
		Sodan: &pb.Sodan{
			Id:        uint64(sodan.ID),
			Title:     sodan.Title,
			Text:      sodan.Text,
			CreaterID: sodan.CreaterID,
			IsClosed:  sodan.IsClosed,
		},
	}
	return res, nil
}

func (s *server) CloseSodan(ctx context.Context, req *pb.CloseSodanRequest) (*empty.Empty, error) {
	return nil, nil
}

func (s *server) CreateReply(ctx context.Context, req *pb.CreateReplyRequest) (*pb.CreateReplyResponse, error) {
	reply := &dto.ReplyDto{
		Text:      req.GetText(),
		CreaterID: req.GetCreaterID(),
		SodanID:   uint(req.GetSodanID()),
	}

	id, err := s.replyService.CreateReply(reply)
	if err != nil {
		return nil, err
	}

	res := &pb.CreateReplyResponse{
		Id: uint64(id),
	}
	return res, nil
}

func (s *server) GetReply(ctx context.Context, req *pb.GetReplyRequest) (*pb.GetReplyResponse, error) {
	id := req.GetId()
	reply, err := s.replyService.FindByID(uint(id))
	if err != nil {
		return nil, err
	}

	res := &pb.GetReplyResponse{
		Reply: &pb.Reply{
			Id:        uint64(reply.ID),
			Text:      reply.Text,
			CreaterID: reply.CreaterID,
			SodanID:   uint64(reply.SodanID),
		},
	}
	return res, nil
}

func (s *server) GetReplies(ctx context.Context, req *pb.GetRepliesRequest) (*pb.GetRepliesResponse, error) {
	return nil, nil
}

func (s *server) SubscribeSodan(req *pb.SubscribeSodanRequest, stream pb.APIService_SubscribeSodanServer) error {
	return nil
}

func (s *server) StopSubscribeSodan(ctx context.Context, req *pb.StopSubscribeSodanRequest) (*empty.Empty, error) {
	return nil, nil
}
