package validate

import (
	"errors"

	apiv1 "github.com/pirosiki197/sodan-grpc/pkg/grpc/pb/api/v1"
)

func ValidateTags(tags []*apiv1.Tag) error {
	for _, t := range tags {
		if t.Name == "" {
			return errors.New("name is empty")
		}
	}
	return nil
}

func ValidateReply(reply *apiv1.Reply) error {
	if reply.Text == "" {
		return errors.New("text is empty")
	}
	if reply.SodanId == 0 {
		return errors.New("sodan_id is empty")
	}
	if reply.CreaterId == "" {
		return errors.New("creater_id is empty")
	}
	return nil
}
