package service

import (
	"github.com/pirosiki197/sodan-grpc/pkg/container"
	apiv1 "github.com/pirosiki197/sodan-grpc/pkg/grpc/pb/api/v1"
	"github.com/pirosiki197/sodan-grpc/pkg/repository/model"
	"github.com/pirosiki197/sodan-grpc/pkg/util/pbconv"
	"github.com/pirosiki197/sodan-grpc/pkg/util/validate"
)

type ReplyService interface {
	FindByID(id uint) (*model.Reply, error)
	FindBySodanID(sodanID uint) ([]*model.Reply, error)
	CreateReply(replyData *apiv1.Reply) (uint, error)
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

func (s *replyService) CreateReply(replyData *apiv1.Reply) (uint, error) {
	if err := validate.ValidateReply(replyData); err != nil {
		return 0, err
	}
	reply := pbconv.ToReplyModel(replyData)
	repo := s.container.GetRepository()

	return repo.CreateReply(reply)
}
