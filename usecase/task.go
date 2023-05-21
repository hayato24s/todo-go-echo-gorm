package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/hayato24s/todo-echo-gorm/apperr"
	"github.com/hayato24s/todo-echo-gorm/entity"
	"github.com/hayato24s/todo-echo-gorm/port"

	"github.com/google/uuid"
)

type CreateTaskIn struct {
	Title string
}

// CreateTask creates a new task.
//
// The following errors may be returned.
//   - ErrUnauthorized
func (uc *UseCase) CreateTask(ctx context.Context, in *CreateTaskIn) (*entity.Task, error) {
	userID, err := uc.Authenticate(ctx)
	if err != nil {
		return nil, err
	}

	task := &entity.Task{
		ID:        uuid.New(),
		UserID:    userID,
		Title:     in.Title,
		Completed: false,
		CreatedAt: time.Now(),
	}
	if err := task.Validate(); err != nil {
		return nil, err
	}

	return task, uc.r.CreateTask(ctx, task)
}

type GetTasksIn struct {
	Completed *bool
	Limit     uint
	Offset    uint // zero-indexed
}

// GetTasks gets the tasks.
//
// The following errors may be returned.
//   - ErrUnauthorized
//   - ErrValidation
func (uc *UseCase) GetTasks(ctx context.Context, in GetTasksIn) (tasks []entity.Task, total uint64, err error) {
	userID, err := uc.Authenticate(ctx)
	if err != nil {
		return
	}

	if in.Limit < 1 || 100 < in.Limit {
		err = fmt.Errorf("%w : Limit must be between 1 and 100", apperr.ErrValidation)
		return
	}

	total, err = uc.r.CountTaskByUserID(ctx, userID)
	if err != nil {
		return
	}

	if total <= uint64(in.Offset) {
		return
	}

	tasks, err = uc.r.FindTasks(ctx, &port.FindTasksConds{
		UserID:    &userID,
		Completed: in.Completed,
		Limit:     &in.Limit,
		Offset:    &in.Offset,
	})
	return
}

type UpdateTaskIn struct {
	ID        uuid.UUID
	Title     *string
	Completed *bool
}

/*
UpdateTask updates the title or completed of the task.

The following errors may be returned.
  - ErrUnauthorized
  - ErrTaskNotFound
  - ErrValidation
*/
func (uc *UseCase) UpdateTask(ctx context.Context, in UpdateTaskIn) error {
	userID, err := uc.Authenticate(ctx)
	if err != nil {
		return err
	}

	task, err := uc.r.FindTaskByIDUserID(ctx, in.ID, userID)
	if err != nil {
		return err
	}

	if in.Title != nil {
		task.Title = *in.Title
		if err := task.ValidateTitle(); err != nil {
			return err
		}
	}

	if in.Completed != nil {
		task.Completed = *in.Completed
	}

	return uc.r.UpdateTaskByID(ctx, task)
}

// DeleteTask deletes the task.
//
// The following errors may be returned.
//   - ErrUnauthorized
//   - ErrTaskNotFound
func (uc *UseCase) DeleteTask(ctx context.Context, id uuid.UUID) error {
	userID, err := uc.Authenticate(ctx)
	if err != nil {
		return err
	}

	if id == uuid.Nil {
		return apperr.ErrValidation
	}

	_, err = uc.r.FindTaskByIDUserID(ctx, id, userID)
	if err != nil {
		return err
	}
	return uc.r.DeleteTaskByID(ctx, id)
}
