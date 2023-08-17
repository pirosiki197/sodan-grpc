package dto

import (
	"errors"
	"slices"

	"github.com/pirosiki197/sodan-grpc/pkg/repository/model"
)

type SodanDto struct {
	Title     string       `json:"title"`
	Text      string       `json:"text"`
	CreaterID string       `json:"creater_id"`
	Tags      []*model.Tag `json:"tags"`
}

func (d *SodanDto) Validate() error {
	if d.Title == "" {
		return errors.New("title is empty")
	}
	if d.CreaterID == "" {
		return errors.New("creater_id is empty")
	}
	d.Tags = slices.DeleteFunc(d.Tags, func(t *model.Tag) bool {
		return t.Name == ""
	})
	return nil
}

func (d *SodanDto) ToSodan() *model.Sodan {
	return &model.Sodan{
		Title:     d.Title,
		Text:      d.Text,
		CreaterID: d.CreaterID,
		Tags:      d.Tags,
	}
}
