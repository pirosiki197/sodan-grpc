package repository

import "github.com/pirosiki197/sodan-grpc/pkg/repository/model"

type SodanRepository interface {
	FindSodanByID(id uint) (*model.Sodan, error)
	CreateSodan(sodan *model.Sodan) (uint, error)
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

func (r *repository) CreateSodan(sodan *model.Sodan) (uint, error) {
	err := r.db.Create(sodan).Error
	if err != nil {
		return 0, err
	}

	return sodan.ID, nil
}

func (r *repository) AddTags(sodanID uint, tags []*model.Tag) error {
	err := r.db.Model(&model.Sodan{ID: sodanID}).Association("Tags").Append(tags)
	if err != nil {
		return err
	}

	return nil
}
