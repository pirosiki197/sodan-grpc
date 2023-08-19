package pbconv

import (
	apiv1 "github.com/pirosiki197/sodan-grpc/pkg/grpc/pb/api/v1"
	"github.com/pirosiki197/sodan-grpc/pkg/repository/model"
)

func ToSodanData(sodan *model.Sodan) *apiv1.Sodan {
	return &apiv1.Sodan{
		Title:     sodan.Title,
		Text:      sodan.Text,
		CreaterId: sodan.CreaterID,
		Tags:      ToTagsData(sodan.Tags),
	}
}

func ToSodansData(sodans []*model.Sodan) []*apiv1.Sodan {
	var result []*apiv1.Sodan
	for _, s := range sodans {
		result = append(result, ToSodanData(s))
	}
	return result
}

func ToTagData(tag *model.Tag) *apiv1.Tag {
	return &apiv1.Tag{
		Name: tag.Name,
	}
}

func ToTagsData(tags []*model.Tag) []*apiv1.Tag {
	var result []*apiv1.Tag
	for _, t := range tags {
		result = append(result, ToTagData(t))
	}
	return result
}

func ToReplyData(reply *model.Reply) *apiv1.Reply {
	return &apiv1.Reply{
		Text:      reply.Text,
		SodanId:   uint64(reply.SodanID),
		CreaterId: reply.CreaterID,
	}
}

func ToRepliesData(replies []*model.Reply) []*apiv1.Reply {
	var result []*apiv1.Reply
	for _, r := range replies {
		result = append(result, ToReplyData(r))
	}
	return result
}
