package port

import (
	"context"

	"github.com/google/uuid"
	"github.com/hayato24s/todo-echo-gorm/entity"
)

type ITaskRepository interface {
	// CreateTask creates a new task.
	CreateTask(ctx context.Context, task *entity.Task) error

	// FindTaskByIDUserID returns the task specified by the given id and userID.
	//
	// The following errors may be returned.
	//   - ErrTaskNotFound.
	FindTaskByIDUserID(ctx context.Context, id, userID uuid.UUID) (*entity.Task, error)

	// FindTasks returns the tasks which satisfy the given conditions in order of newest to oldest.
	FindTasks(ctx context.Context, conds *FindTasksConds) ([]entity.Task, error)

	// CountTaskByUserID returns the total number of the tasks specified by userID.
	CountTaskByUserID(ctx context.Context, userID uuid.UUID) (uint64, error)

	// UpdateTaskByID updates the following fields of the given task.
	//   - Title
	//   - Completed
	UpdateTaskByID(ctx context.Context, task *entity.Task) error

	// DeleteTaskByID deletes the task specified by the given id.
	DeleteTaskByID(ctx context.Context, id uuid.UUID) error

	// DeleteTasksByUserID deletes the tasks specified by the given userID.
	DeleteTasksByUserID(ctx context.Context, userID uuid.UUID) error
}

type FindTasksConds struct {
	UserID    *uuid.UUID
	Completed *bool
	Limit     *uint
	Offset    *uint
}
