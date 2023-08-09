package repository

import "github.com/pirosiki197/sodan-grpc/pkg/repository/model"

type ReplyRepository interface {
	FindReplyByID(id uint) (*model.Reply, error)
	FindRepliesBySodanID(sodanID uint) ([]*model.Reply, error)
	CreateReply(reply *model.Reply) (uint, error)
}

func (r *repository) FindReplyByID(id uint) (*model.Reply, error) {
	reply := new(model.Reply)
	err := r.db.First(reply, id).Error
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (r *repository) FindRepliesBySodanID(sodanID uint) ([]*model.Reply, error) {
	replies := make([]*model.Reply, 0)
	err := r.db.Where("sodan_id = ?", sodanID).Find(&replies).Error
	if err != nil {
		return nil, err
	}

	return replies, nil
}

func (r *repository) CreateReply(reply *model.Reply) (uint, error) {
	err := r.db.Create(reply).Error
	if err != nil {
		return 0, err
	}

	return reply.ID, nil
}
