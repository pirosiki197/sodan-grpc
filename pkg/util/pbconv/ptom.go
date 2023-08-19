package pbconv

import (
	apiv1 "github.com/pirosiki197/sodan-grpc/pkg/grpc/pb/api/v1"
	"github.com/pirosiki197/sodan-grpc/pkg/repository/model"
)

func ToTagModel(t *apiv1.Tag) *model.Tag {
	return &model.Tag{
		Name: t.Name,
	}
}

func ToTagModels(tags []*apiv1.Tag) []*model.Tag {
	var result []*model.Tag
	for _, t := range tags {
		result = append(result, ToTagModel(t))
	}
	return result
}

func ToSodanModel(sodan *apiv1.Sodan) *model.Sodan {
	return &model.Sodan{
		Title:     sodan.Title,
		Text:      sodan.Text,
		CreaterID: sodan.CreaterId,
		Tags:      ToTagModels(sodan.Tags),
	}
}

func ToReplyModel(reply *apiv1.Reply) *model.Reply {
	return &model.Reply{
		Text:      reply.Text,
		SodanID:   uint(reply.SodanId),
		CreaterID: reply.CreaterId,
	}
}
