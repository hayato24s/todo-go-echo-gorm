package repository_test

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hayato24s/todo-echo-gorm/apperr"
	"github.com/hayato24s/todo-echo-gorm/entity"
)

func FixtureUser(suite *RepositoryTestSuite, save bool) *entity.User {
	user := &entity.User{
		ID:       uuid.New(),
		Name:     fmt.Sprintf("user-%s", time.Now().Format(time.RFC3339Nano)),
		Password: "password",
	}

	if save {
		if err := suite.r.CreateUser(context.Background(), user); err != nil {
			suite.T().Fatal(err)
		}
	}
	return user
}

func (suite *RepositoryTestSuite) TestCreateUesr() {
	user := FixtureUser(suite, false)
	suite.Run("normal", func() {
		err := suite.r.CreateUser(context.Background(), user)
		if err != nil {
			suite.T().Fatal(err)
		}
	})
}

func (suite *RepositoryTestSuite) TestFindUserByName() {
	user := FixtureUser(suite, true)

	suite.Run("normal", func() {
		got, err := suite.r.FindUserByName(context.Background(), user.Name)
		if err != nil {
			suite.T().Fatal(err)
		}
		if user.ID != got.ID {
			suite.T().Fatal("id should be same")
		}
	})

	suite.Run("not found", func() {
		_, err := suite.r.FindUserByName(context.Background(), uuid.NewString())
		if !errors.Is(err, apperr.ErrUserNotFound) {
			suite.T().Fatal("user should be not found")
		}
	})
}

func (suite *RepositoryTestSuite) TestDeleteUserByID() {
	user := FixtureUser(suite, true)

	suite.Run("normal", func() {
		err := suite.r.DeleteUserByID(context.Background(), user.ID)
		if err != nil {
			suite.T().Fatal(err)
		}
		_, err = suite.r.FindUserByName(context.Background(), user.Name)
		if err == nil {
			suite.T().Fatal("user should be not found")
		}
		if !errors.Is(err, apperr.ErrUserNotFound) {
			suite.T().Fatal(err)
		}
	})
}
