package repository

import (
	"errors"

	"github.com/pirosiki197/sodan-grpc/pkg/repository/model"
)

type SodanRepository interface {
	FindSodanByID(id uint) (*model.Sodan, error)
	FindSodans(where ...any) ([]*model.Sodan, error)
	FindSodansByTag(tag string) ([]*model.Sodan, error)
	CreateSodan(sodan *model.Sodan) (uint, error)
	UpdateSodan(sodan *model.Sodan) error
	AddTags(sodanID uint, tags []*model.Tag) error
}

func (r *repository) FindSodanByID(id uint) (*model.Sodan, error) {
	sodan := new(model.Sodan)
	err := r.db.Preload("Tags").First(sodan, id).Error
	if err != nil {
		return nil, err
	}

	return sodan, nil
}

func (r *repository) FindSodans(where ...any) ([]*model.Sodan, error) {
	sodans := make([]*model.Sodan, 0)
	err := r.db.Preload("Tags").Order("created_at desc").Find(&sodans, where...).Error
	if err != nil {
		return nil, err
	}

	return sodans, nil
}

func (r *repository) FindSodansByTag(tag string) ([]*model.Sodan, error) {
	sodans := make([]*model.Sodan, 0)
	err := r.db.Debug().Where("id IN (?)", r.db.Table("tags").Where("name = ?", tag).Select("sodan_id")).Preload("Tags").Order("created_at desc").Find(&sodans).Error
	if err != nil {
		return nil, err
	}

	return sodans, nil
}

func (r *repository) CreateSodan(sodan *model.Sodan) (uint, error) {
	err := r.db.Create(sodan).Error
	if err != nil {
		return 0, err
	}

	return sodan.ID, nil
}

func (r *repository) UpdateSodan(sodan *model.Sodan) error {
	if sodan.ID == 0 {
		return errors.New("sodan id is not set")
	}
	result := r.db.Save(sodan)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repository) AddTags(sodanID uint, tags []*model.Tag) error {
	err := r.db.Model(&model.Sodan{ID: sodanID}).Association("Tags").Append(tags)
	if err != nil {
		return err
	}

	return nil
}
