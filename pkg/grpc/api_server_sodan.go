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
	sodan := &dto.SodanDto{
		Title:     req.Msg.Title,
		Text:      req.Msg.Text,
		CreaterID: req.Msg.CreaterId,
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
	req.Header().Set("Access-Control-Allow-Origin", "*")
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
