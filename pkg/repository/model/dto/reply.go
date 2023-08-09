package dto

import (
	"errors"

	"github.com/pirosiki197/sodan-grpc/pkg/repository/model"
)

type ReplyDto struct {
	Text      string `json:"text"`
	SodanID   uint   `json:"sodan_id"`
	CreaterID string `json:"creater_id"`
}

func (d *ReplyDto) Validate() error {
	if d.Text == "" {
		return errors.New("text is empty")
	}
	if d.SodanID == 0 {
		return errors.New("sodan_id is empty")
	}
	if d.CreaterID == "" {
		return errors.New("creater_id is empty")
	}
	return nil
}

func (d *ReplyDto) ToReply() *model.Reply {
	return &model.Reply{
		Text:      d.Text,
		SodanID:   d.SodanID,
		CreaterID: d.CreaterID,
	}
}
