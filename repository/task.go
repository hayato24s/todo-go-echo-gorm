package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/hayato24s/todo-echo-gorm/apperr"
	"github.com/hayato24s/todo-echo-gorm/entity"
	"github.com/hayato24s/todo-echo-gorm/port"
	"gorm.io/gorm"
)

func (r *Repository) CreateTask(ctx context.Context, task *entity.Task) error {
	return r.db.WithContext(ctx).
		Create(task).
		Error
}

func (r *Repository) FindTaskByIDUserID(ctx context.Context, id, userID uuid.UUID) (*entity.Task, error) {
	var task entity.Task
	err := r.db.WithContext(ctx).
		Where("id = ? AND user_id = ?", id, userID).
		Take(&task).
		Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperr.ErrTaskNotFound
	}
	return &task, err
}

func (r *Repository) FindTasks(ctx context.Context, conds *port.FindTasksConds) ([]entity.Task, error) {
	db := r.db.WithContext(ctx)
	if conds.UserID != nil {
		db = db.Where("user_id = ?", *conds.UserID)
	}
	if conds.Completed != nil {
		db = db.Where("completed = ?", *conds.Completed)
	}
	if conds.Limit != nil {
		db = db.Limit(int(*conds.Limit))
	}
	if conds.Offset != nil {
		db = db.Offset(int(*conds.Offset))
	}

	var tasks []entity.Task
	err := db.
		Order("created_at DESC").
		Find(&tasks).
		Error
	return tasks, err
}

func (r *Repository) CountTaskByUserID(ctx context.Context, userID uuid.UUID) (uint64, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Count(&count).
		Error
	return uint64(count), err
}

func (r *Repository) UpdateTaskByID(ctx context.Context, task *entity.Task) error {
	return r.db.WithContext(ctx).
		Model(task).
		Select("title", "completed").
		UpdateColumns(task).
		Error
}

func (r *Repository) DeleteTaskByID(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&entity.Task{}).
		Error
}

func (r *Repository) DeleteTasksByUserID(ctx context.Context, userID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Delete(&entity.Task{}).
		Error
}
