package grpc

import (
	"context"

	"connectrpc.com/connect"
	apiv1 "github.com/pirosiki197/sodan-grpc/pkg/grpc/pb/api/v1"
	"github.com/pirosiki197/sodan-grpc/pkg/repository/model"
	"github.com/pirosiki197/sodan-grpc/pkg/repository/model/dto"
	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) CreateSodan(ctx context.Context, req *connect.Request[apiv1.CreateSodanRequest]) (*connect.Response[apiv1.CreateSodanResponse], error) {
	s.logger.Info("CreateSodan", "req", req.Msg)
	sodan := &dto.SodanDto{
		Title:     req.Msg.Title,
		Text:      req.Msg.Text,
		CreaterID: req.Msg.CreaterId,
		Tags:      lo.Map(req.Msg.Tags, func(t *apiv1.Tag, _ int) *model.Tag { return &model.Tag{Name: t.Name} }),
	}

	id, err := s.sodanService.CreateSodan(sodan)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&apiv1.CreateSodanResponse{
		Id: uint64(id),
	})
	return res, nil
}

func (s *server) GetSodan(ctx context.Context, req *connect.Request[apiv1.GetSodanRequest]) (*connect.Response[apiv1.GetSodanResponse], error) {
	s.logger.Info("GetSodan", "req", req.Msg)
	id := req.Msg.Id
	sodan, err := s.sodanService.FindByID(uint(id))
	if err != nil {
		return nil, err
	}
	// TODO: sodanをapiv1.Sodanに変換する関数を作成する -> pbconv
	res := connect.NewResponse(&apiv1.GetSodanResponse{
		Sodan: &apiv1.Sodan{
			Id:        uint64(sodan.ID),
			Title:     sodan.Title,
			Text:      sodan.Text,
			CreaterId: sodan.CreaterID,
			Tags:      lo.Map(sodan.Tags, func(t *model.Tag, _ int) *apiv1.Tag { return &apiv1.Tag{Name: t.Name} }),
		},
	})
	return res, nil
}

func (s *server) GetSodanList(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[apiv1.GetSodanListResponse], error) {
	sodans, err := s.sodanService.GetSodanList()
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&apiv1.GetSodanListResponse{
		Sodans: lo.Map(sodans, func(s *model.Sodan, _ int) *apiv1.Sodan {
			return &apiv1.Sodan{
				Id:        uint64(s.ID),
				Title:     s.Title,
				Text:      s.Text,
				CreaterId: s.CreaterID,
				Tags:      lo.Map(s.Tags, func(t *model.Tag, _ int) *apiv1.Tag { return &apiv1.Tag{Name: t.Name} }),
			}
		}),
	})
	return res, nil
}

func (s *server) GetSodansByTag(ctx context.Context, req *connect.Request[apiv1.GetSodansByTagRequest]) (*connect.Response[apiv1.GetSodansByTagResponse], error) {
	s.logger.Info("GetSodansByTag", "req", req.Msg)
	tag := req.Msg.TagName
	sodans, err := s.sodanService.FindByTag(tag)
	if err != nil {
		return nil, err
	}

	res := connect.NewResponse(&apiv1.GetSodansByTagResponse{
		Sodans: lo.Map(sodans, func(s *model.Sodan, _ int) *apiv1.Sodan {
			return &apiv1.Sodan{
				Id:        uint64(s.ID),
				Title:     s.Title,
				Text:      s.Text,
				CreaterId: s.CreaterID,
				Tags:      lo.Map(s.Tags, func(t *model.Tag, _ int) *apiv1.Tag { return &apiv1.Tag{Name: t.Name} }),
			}
		}),
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
