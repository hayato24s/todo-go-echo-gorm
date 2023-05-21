package repository_test

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/hayato24s/todo-echo-gorm/apperr"
	"github.com/hayato24s/todo-echo-gorm/entity"
	"github.com/hayato24s/todo-echo-gorm/port"
)

func FixtureTask(suite *RepositoryTestSuite, save bool, userID uuid.UUID) *entity.Task {
	titles := []string{
		"read documentation",
		"deploy app",
		"write test code",
		"meeting with clients",
		"make a wireframe",
	}
	task := &entity.Task{
		ID:        uuid.New(),
		UserID:    userID,
		Title:     titles[rand.Intn(len(titles))],
		Completed: rand.Intn(2) == 1,
		CreatedAt: time.Now(),
	}

	if save {
		if err := suite.r.CreateTask(context.Background(), task); err != nil {
			suite.T().Fatal(err)
		}
	}
	return task
}

func FixtureTasks(suite *RepositoryTestSuite, save bool, userID uuid.UUID, num int) []entity.Task {
	var tasks []entity.Task
	for i := 0; i < num; i++ {
		task := FixtureTask(suite, save, userID)
		tasks = append(tasks, *task)
	}
	return tasks
}

func (suite *RepositoryTestSuite) TestCreateTask() {
	user := FixtureUser(suite, true)
	task := FixtureTask(suite, false, user.ID)

	suite.Run("normal", func() {
		err := suite.r.CreateTask(context.Background(), task)
		if err != nil {
			suite.T().Fatal(err)
		}
	})
}

func (suite *RepositoryTestSuite) TestFindTaskByIDUserID() {
	user := FixtureUser(suite, true)
	task := FixtureTask(suite, true, user.ID)

	suite.Run("normal", func() {
		_, err := suite.r.FindTaskByIDUserID(context.Background(), task.ID, user.ID)
		if err != nil {
			suite.T().Fatal(err)
		}
	})

	suite.Run("not found 1", func() {
		_, err := suite.r.FindTaskByIDUserID(context.Background(), uuid.New(), user.ID)
		if !errors.Is(err, apperr.ErrTaskNotFound) {
			suite.T().Fatal("task should be not found")
		}
	})

	suite.Run("not found 2", func() {
		_, err := suite.r.FindTaskByIDUserID(context.Background(), task.ID, uuid.New())
		if !errors.Is(err, apperr.ErrTaskNotFound) {
			suite.T().Fatal("task should be not found")
		}
	})
}

func (suite *RepositoryTestSuite) TestFindTasks() {
	user := FixtureUser(suite, true)
	tasks := FixtureTasks(suite, true, user.ID, 10)

	suite.Run("UserID", func() {
		conds := &port.FindTasksConds{
			UserID: &user.ID,
		}
		got, err := suite.r.FindTasks(context.Background(), conds)
		if err != nil {
			suite.T().Fatal(err)
		}
		if len(tasks) != len(got) {
			suite.T().Fatal("got should be same length")
		}
	})

	suite.Run("UserID Limit Offset", func() {
		conds := &port.FindTasksConds{
			UserID: &user.ID,
			Limit:  new(uint),
			Offset: new(uint),
		}
		*conds.Limit = 4
		*conds.Offset = 2

		got, err := suite.r.FindTasks(context.Background(), conds)
		if err != nil {
			suite.T().Fatal(err)
		}
		if len(got) != int(*conds.Limit) {
			suite.T().Fatal("invalid length")
		}
	})
}

func (suite *RepositoryTestSuite) TestUpdateTaskByID() {
	user := FixtureUser(suite, true)
	task := FixtureTask(suite, true, user.ID)

	suite.Run("normal", func() {
		task.Title = "updated title"
		task.Completed = !task.Completed

		err := suite.r.UpdateTaskByID(context.Background(), task)
		if err != nil {
			suite.T().Fatal(err)
		}
		got, err := suite.r.FindTaskByIDUserID(context.Background(), task.ID, task.UserID)
		if err != nil {
			suite.T().Fatal(err)
		}
		if task.Title != got.Title || task.Completed != got.Completed {
			suite.T().Fatal("task should be updated")
		}
	})
}

func (suite *RepositoryTestSuite) TestDeleteTaskByID() {
	user := FixtureUser(suite, true)
	task := FixtureTask(suite, true, user.ID)

	suite.Run("normal", func() {
		err := suite.r.DeleteTaskByID(context.Background(), task.ID)
		if err != nil {
			suite.T().Fatal(err)
		}
		_, err = suite.r.FindTaskByIDUserID(context.Background(), task.ID, task.UserID)
		if !errors.Is(err, apperr.ErrTaskNotFound) {
			suite.T().Fatal("task should not be found")
		}
	})
}

func (suite *RepositoryTestSuite) TestDeleteTasksByUserID() {
	user := FixtureUser(suite, true)
	_ = FixtureTasks(suite, true, user.ID, 10)

	suite.Run("normal", func() {
		err := suite.r.DeleteTasksByUserID(context.Background(), user.ID)
		if err != nil {
			suite.T().Fatal(err)
		}
		conds := &port.FindTasksConds{
			UserID: &user.ID,
		}
		got, err := suite.r.FindTasks(context.Background(), conds)
		if err != nil {
			suite.T().Fatal(err)
		}
		if len(got) != 0 {
			suite.T().Fatal("tasks should be deleted")
		}
	})
}
