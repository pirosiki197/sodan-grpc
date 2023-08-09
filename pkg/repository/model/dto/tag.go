package dto

import (
	"errors"

	"github.com/pirosiki197/sodan-grpc/pkg/repository/model"
)

type Tag struct {
	Name string `json:"name"`
}

func (t *Tag) Validate() error {
	if t.Name == "" {
		return errors.New("name is empty")
	}

	return nil
}

func (t *Tag) ToTag() *model.Tag {
	return &model.Tag{
		Name: t.Name,
	}
}
