package service

import (
	"github.com/pirosiki197/sodan-grpc/pkg/container"
	"github.com/pirosiki197/sodan-grpc/pkg/repository/model"
	"github.com/pirosiki197/sodan-grpc/pkg/repository/model/dto"
)

type ReplyService interface {
	FindByID(id uint) (*model.Reply, error)
	FindBySodanID(sodanID uint) ([]*model.Reply, error)
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

func (s *replyService) FindBySodanID(sodanID uint) ([]*model.Reply, error) {
	repo := s.container.GetRepository()

	return repo.FindRepliesBySodanID(sodanID)
}

func (s *replyService) CreateReply(dto *dto.ReplyDto) (uint, error) {
	if err := dto.Validate(); err != nil {
		return 0, err
	}
	reply := dto.ToReply()
	repo := s.container.GetRepository()

	return repo.CreateReply(reply)
}
