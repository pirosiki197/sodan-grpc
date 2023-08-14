package service

import (
	"github.com/pirosiki197/sodan-grpc/pkg/container"
	"github.com/pirosiki197/sodan-grpc/pkg/repository/model"
	"github.com/pirosiki197/sodan-grpc/pkg/repository/model/dto"
	"github.com/pirosiki197/sodan-grpc/pkg/util"
	"github.com/samber/lo"
)

type SodanService interface {
	FindByID(id uint) (*model.Sodan, error)
	GetSodanList() ([]*model.Sodan, error)
	FindByTag(tag string) ([]*model.Sodan, error)
	CreateSodan(dto *dto.SodanDto) (uint, error)
	CloseSodan(id uint) error
	AddTags(sodanID string, dto []*dto.Tag) error
}

type sodanService struct {
	container container.Container
}

func NewSodanService(container container.Container) SodanService {
	return &sodanService{container: container}
}

func (s *sodanService) FindByID(id uint) (*model.Sodan, error) {
	repo := s.container.GetRepository()
	sodan, err := repo.FindSodanByID(id)
	if err != nil {
		return nil, err
	}

	return sodan, nil
}

func (s *sodanService) GetSodanList() ([]*model.Sodan, error) {
	repo := s.container.GetRepository()
	sodans, err := repo.FindSodans()
	if err != nil {
		return nil, err
	}

	return sodans, nil
}

func (s *sodanService) FindByTag(tag string) ([]*model.Sodan, error) {
	repo := s.container.GetRepository()
	sodans, err := repo.FindSodansByTag(tag)
	if err != nil {
		return nil, err
	}

	return sodans, nil
}

func (s *sodanService) CreateSodan(dto *dto.SodanDto) (uint, error) {
	if err := dto.Validate(); err != nil {
		return 0, err
	}
	sodan := dto.ToSodan()
	sodan.Tags = deleteDuplicateTags(sodan.Tags)
	repo := s.container.GetRepository()

	return repo.CreateSodan(sodan)
}

func (s *sodanService) CloseSodan(id uint) error {
	repo := s.container.GetRepository()
	closedSodan := model.Sodan{ID: id, IsClosed: true}
	return repo.UpdateSodan(&closedSodan)
}

// AddTags does not add tags that already exist.
func (s *sodanService) AddTags(sodanID string, dto []*dto.Tag) error {
	var tags []*model.Tag
	for _, tag := range dto {
		if err := tag.Validate(); err != nil {
			return err
		}
		tags = append(tags, tag.ToTag())
	}
	repo := s.container.GetRepository()
	sodan, err := repo.FindSodanByID(util.ConvertToUint(sodanID))
	if err != nil {
		return err
	}
	tags = deleteDuplicateTags(tags)
	filteredTags := fileterTags(tags, sodan)
	if len(filteredTags) == 0 {
		return nil
	}

	return repo.AddTags(util.ConvertToUint(sodanID), filteredTags)
}

func deleteDuplicateTags(tags []*model.Tag) []*model.Tag {
	tagsMap := make(map[string]*model.Tag)
	for _, tag := range tags {
		if _, ok := tagsMap[tag.Name]; ok {
			continue
		}
		tagsMap[tag.Name] = tag
	}
	return lo.MapToSlice(tagsMap, func(_ string, v *model.Tag) *model.Tag { return v })
}

func fileterTags(tags []*model.Tag, sodan *model.Sodan) []*model.Tag {
	filteredTags := lo.Filter(tags, func(tag *model.Tag, _ int) bool {
		for _, t := range sodan.Tags {
			if tag.Name == t.Name {
				return false
			}
		}
		return true
	})
	return filteredTags
}
