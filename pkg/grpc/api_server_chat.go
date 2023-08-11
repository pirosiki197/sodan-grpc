package grpc

import (
	"context"
	"slices"

	"connectrpc.com/connect"
	apiv1 "github.com/pirosiki197/sodan-grpc/pkg/grpc/pb/api/v1"
	"github.com/pirosiki197/sodan-grpc/pkg/repository/model"
	"github.com/pirosiki197/sodan-grpc/pkg/repository/model/dto"
	"github.com/samber/lo"
)

type newReplyInfo struct {
	id      uint64
	sodanID uint64
}

var (
	// 新しいリプライの情報を受け取るチャンネル
	chs []chan<- newReplyInfo
	// 新しいリプライのSOdanIDを受け取るチャンネル
	newReplych chan newReplyInfo = make(chan newReplyInfo, 1)
)

func (s *server) CreateReply(ctx context.Context, req *connect.Request[apiv1.CreateReplyRequest]) (*connect.Response[apiv1.CreateReplyResponse], error) {
	reply := &dto.ReplyDto{
		Text:      req.Msg.GetText(),
		CreaterID: req.Msg.GetCreaterId(),
		SodanID:   uint(req.Msg.GetSodanId()),
	}

	id, err := s.replyService.CreateReply(reply)
	if err != nil {
		return nil, err
	}

	go func() {
		newReplych <- newReplyInfo{id: uint64(id), sodanID: req.Msg.GetSodanId()}
	}()

	res := connect.NewResponse(&apiv1.CreateReplyResponse{
		Id: uint64(id),
	})
	return res, nil
}

func (s *server) GetReply(ctx context.Context, req *connect.Request[apiv1.GetReplyRequest]) (*connect.Response[apiv1.GetReplyResponse], error) {
	id := req.Msg.GetId()
	reply, err := s.replyService.FindByID(uint(id))
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&apiv1.GetReplyResponse{
		Reply: &apiv1.Reply{
			Id:        uint64(reply.ID),
			Text:      reply.Text,
			CreaterId: reply.CreaterID,
			SodanId:   uint64(reply.SodanID),
		},
	})
	return res, nil
}

func (s *server) GetReplies(ctx context.Context, req *connect.Request[apiv1.GetRepliesRequest]) (*connect.Response[apiv1.GetRepliesResponse], error) {
	sodanID := req.Msg.SodanId
	replies, err := s.replyService.FindBySodanID(uint(sodanID))
	if err != nil {
		return nil, err
	}
	// TODO: pbconvに処理を移す
	res := connect.NewResponse(&apiv1.GetRepliesResponse{
		Replies: lo.Map(replies, func(reply *model.Reply, _ int) *apiv1.Reply {
			return &apiv1.Reply{
				Id:        uint64(reply.ID),
				Text:      reply.Text,
				CreaterId: reply.CreaterID,
				SodanId:   uint64(reply.SodanID),
			}
		}),
	})
	return res, nil
}

func (s *server) SubscribeSodan(ctx context.Context, req *connect.Request[apiv1.SubscribeSodanRequest], stream *connect.ServerStream[apiv1.SubscribeSodanResponse]) error {
	sodanId := req.Msg.Id
	ch := make(chan newReplyInfo)
	appendCh(ch)
	defer removeCh(ch)
	for {
		select {
		case newReply := <-ch:
			if newReply.sodanID == sodanId {
				reply, err := s.replyService.FindByID(uint(newReply.id))
				if err != nil {
					return err
				}
				res := &apiv1.SubscribeSodanResponse{
					Reply: &apiv1.Reply{
						Id:        uint64(reply.ID),
						Text:      reply.Text,
						CreaterId: reply.CreaterID,
						SodanId:   uint64(reply.SodanID),
					},
				}
				if err := stream.Send(res); err != nil {
					s.logger.Error("stream send error", "err", err)
					return err
				}
			}
		case <-ctx.Done():
			return nil
		}
	}
}

func appendCh(ch chan<- newReplyInfo) {
	chs = append(chs, ch)
}

func removeCh(ch chan<- newReplyInfo) {
	slices.DeleteFunc(chs, func(c chan<- newReplyInfo) bool {
		return c == ch
	})
}

// お試し
func checkNewReply() {
	for {
		newID := <-newReplych
		for _, ch := range chs {
			ch <- newID
		}
	}
}