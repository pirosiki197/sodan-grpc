package service

import (
	"errors"

	"github.com/pirosiki197/sodan-grpc/pkg/container"
	"github.com/pirosiki197/sodan-grpc/pkg/repository/model"
	"github.com/pirosiki197/sodan-grpc/pkg/repository/model/dto"
	"github.com/pirosiki197/sodan-grpc/pkg/util"
)

type ReplyService interface {
	FindByID(id uint) (*model.Reply, error)
	FindBySodanID(sodanID string) ([]*model.Reply, error)
	CreateReply(dto *dto.ReplyDto) (uint, error)
}

type replyService struct {
	container container.Container
}

func NewReplyService(container container.Container) *replyService {
	return &replyService{container: container}
}

func (s *replyService) FindByID(id uint) (*model.Reply, error) {
	repo := s.container.GetRepository()

	return repo.FindReplyByID(id)
}

func (s *replyService) FindBySodanID(sodanID string) ([]*model.Reply, error) {
	if sodanID == "" {
		return nil, errors.New("sodanID is empty")
	}
	repo := s.container.GetRepository()

	return repo.FindRepliesBySodanID(util.ConvertToUint(sodanID))
}

func (s *replyService) CreateReply(dto *dto.ReplyDto) (uint, error) {
	if err := dto.Validate(); err != nil {
		return 0, err
	}
	reply := dto.ToReply()
	repo := s.container.GetRepository()

	return repo.CreateReply(reply)
}
