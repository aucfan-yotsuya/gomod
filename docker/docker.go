package docker

import (
	"github.com/ory/dockertest/v3"
	"gorm.io/gorm"
)

type (
	Docker struct {
		Target []*Target
	}
	Target struct {
		Pool     *dockertest.Pool
		Resource *dockertest.Resource
		Gorm     *gorm.DB
	}
	Opt struct {
		Endpoint         string
		DockerRunOptions *dockertest.RunOptions
	}
)

func New() *Docker { return new(Docker) }
