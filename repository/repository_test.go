package repository_test

import (
	"testing"

	"github.com/hayato24s/todo-echo-gorm/port"
	"github.com/hayato24s/todo-echo-gorm/repository"
	"github.com/stretchr/testify/suite"
)

type RepositoryTestSuite struct {
	suite.Suite
	r port.IRepository
}

func (suite *RepositoryTestSuite) SetupSuite() {
	var err error
	suite.r, err = repository.NewRepository()
	if err != nil {
		panic(err)
	}
}

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}
