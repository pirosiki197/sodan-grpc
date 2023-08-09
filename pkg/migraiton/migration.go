package migration

import (
	"github.com/pirosiki197/sodan-grpc/pkg/container"
	"github.com/pirosiki197/sodan-grpc/pkg/repository/model"
)

func CreateDB(container container.Container) {
	repo := container.GetRepository()

	_ = repo.AutoMigrate(&model.Sodan{})
	_ = repo.AutoMigrate(&model.Reply{})
	_ = repo.AutoMigrate(&model.Tag{})
}
