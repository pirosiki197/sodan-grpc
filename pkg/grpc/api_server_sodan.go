package grpc

import (
	"context"

	"connectrpc.com/connect"
	apiv1 "github.com/pirosiki197/sodan-grpc/pkg/grpc/pb/api/v1"
	"github.com/pirosiki197/sodan-grpc/pkg/util/pbconv"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) CreateSodan(ctx context.Context, req *connect.Request[apiv1.CreateSodanRequest]) (*connect.Response[apiv1.CreateSodanResponse], error) {
	sodan := &apiv1.Sodan{
		Title:     req.Msg.Title,
		Text:      req.Msg.Text,
		CreaterId: req.Msg.CreaterId,
		Tags:      req.Msg.Tags,
	}

	id, err := s.sodanService.CreateSodan(sodan)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&apiv1.CreateSodanResponse{
		Id: uint32(id),
	})
	return res, nil
}

func (s *server) GetSodan(ctx context.Context, req *connect.Request[apiv1.GetSodanRequest]) (*connect.Response[apiv1.GetSodanResponse], error) {
	id := req.Msg.Id
	sodan, err := s.sodanService.FindByID(uint(id))
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&apiv1.GetSodanResponse{
		Sodan: pbconv.ToSodanData(sodan),
	})
	return res, nil
}

func (s *server) GetSodanList(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[apiv1.GetSodanListResponse], error) {
	sodans, err := s.sodanService.GetSodanList()
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&apiv1.GetSodanListResponse{
		Sodans: pbconv.ToSodansData(sodans),
	})
	return res, nil
}

func (s *server) GetSodansByTag(ctx context.Context, req *connect.Request[apiv1.GetSodansByTagRequest]) (*connect.Response[apiv1.GetSodansByTagResponse], error) {
	tag := req.Msg.TagName
	sodans, err := s.sodanService.FindByTag(tag)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&apiv1.GetSodansByTagResponse{
		Sodans: pbconv.ToSodansData(sodans),
	})
	return res, nil
}

func (s *server) CloseSodan(ctx context.Context, req *connect.Request[apiv1.CloseSodanRequest]) (*connect.Response[emptypb.Empty], error) {
	id := req.Msg.Id
	err := s.sodanService.CloseSodan(uint(id))
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&emptypb.Empty{})
	return res, nil
}
