package grpc

import (
	"context"
	"fmt"
	"sync"

	"connectrpc.com/connect"
	apiv1 "github.com/pirosiki197/sodan-grpc/pkg/grpc/pb/api/v1"
	"github.com/pirosiki197/sodan-grpc/pkg/repository/model"
	"github.com/pirosiki197/sodan-grpc/pkg/util/pbconv"
)

type newReplyInfo struct {
	id      uint32
	sodanID uint32
}

type subscriber struct {
	m   sync.Mutex
	chs []chan<- *model.Reply
}

var (
	//
	subsc subscriber
	// 新しいリプライのSodanIDを受け取るチャンネル
	newReplych chan newReplyInfo = make(chan newReplyInfo)
)

func (s *server) CreateReply(ctx context.Context, req *connect.Request[apiv1.CreateReplyRequest]) (*connect.Response[apiv1.CreateReplyResponse], error) {
	s.logger.Info("CreateReply", "req", req.Msg)

	// sodanが存在するか確認
	_, err := s.sodanService.FindByID(uint(req.Msg.SodanId))
	if err != nil {
		return nil, err
	}

	reply := &apiv1.Reply{
		Text:      req.Msg.Text,
		CreaterId: req.Msg.CreaterId,
		SodanId:   req.Msg.SodanId,
	}

	id, err := s.replyService.CreateReply(reply)
	if err != nil {
		return nil, err
	}

	go func() {
		newReplych <- newReplyInfo{id: uint32(id), sodanID: req.Msg.GetSodanId()}
	}()

	res := connect.NewResponse(&apiv1.CreateReplyResponse{
		Id: uint32(id),
	})
	return res, nil
}

func (s *server) GetReply(ctx context.Context, req *connect.Request[apiv1.GetReplyRequest]) (*connect.Response[apiv1.GetReplyResponse], error) {
	id := req.Msg.Id
	reply, err := s.replyService.FindByID(uint(id))
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&apiv1.GetReplyResponse{
		Reply: pbconv.ToReplyData(reply),
	})
	return res, nil
}

func (s *server) GetReplies(ctx context.Context, req *connect.Request[apiv1.GetRepliesRequest]) (*connect.Response[apiv1.GetRepliesResponse], error) {
	s.logger.Info("GetReplies", "req", req.Msg)
	sodanID := req.Msg.SodanId
	replies, err := s.replyService.FindBySodanID(uint(sodanID))
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&apiv1.GetRepliesResponse{
		Replies: pbconv.ToRepliesData(replies),
	})
	return res, nil
}

func (s *server) SubscribeSodan(ctx context.Context, req *connect.Request[apiv1.SubscribeSodanRequest], stream *connect.ServerStream[apiv1.SubscribeSodanResponse]) error {
	s.logger.Info("SubscribeSodan", "req", req.Msg)
	sodanId := req.Msg.Id
	ch := make(chan *model.Reply)
	appendCh(ch)
	defer removeCh(ch)
	for {
		select {
		case reply := <-ch:
			if reply.SodanID == uint(sodanId) {
				res := &apiv1.SubscribeSodanResponse{
					Reply: pbconv.ToReplyData(reply),
				}
				if err := stream.Send(res); err != nil {
					s.logger.Error("stream send error", "err", err)
					return err
				}
			}
		case <-ctx.Done():
			s.logger.Info("SubscribeSodan", "ctx done", ctx.Err())
			return nil
		}
	}
}

func appendCh(ch chan<- *model.Reply) {
	subsc.m.Lock()
	defer subsc.m.Unlock()
	subsc.chs = append(subsc.chs, ch)
}

func removeCh(ch chan<- *model.Reply) {
	subsc.m.Lock()
	defer subsc.m.Unlock()
	for i, c := range subsc.chs {
		if c == ch {
			subsc.chs = append(subsc.chs[:i], subsc.chs[i+1:]...)
			break
		}
	}
}

func (s *server) checkNewReply() {
	for {
		newID := <-newReplych
		reply, err := s.replyService.FindByID(uint(newID.id))
		if err != nil {
			fmt.Println("err", err)
			continue
		}
		for _, ch := range subsc.chs {
			ch <- reply
		}
	}
}
