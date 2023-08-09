package container

import (
	"github.com/pirosiki197/sodan-grpc/pkg/config"
	"github.com/pirosiki197/sodan-grpc/pkg/repository"
)

type Container interface {
	GetRepository() repository.Repository
	GetConfig() *config.Config
}

type container struct {
	repo   repository.Repository
	config *config.Config
}

func NewContainer(repo repository.Repository, config *config.Config) Container {
	return &container{repo: repo, config: config}
}

func (c *container) GetRepository() repository.Repository {
	return c.repo
}

func (c *container) GetConfig() *config.Config {
	return c.config
}
