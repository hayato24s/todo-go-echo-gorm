package repository

import (
	"fmt"
	"os"

	"github.com/hayato24s/todo-echo-gorm/port"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func (r *Repository) Begin() (port.IRepository, error) {
	db := r.db.Begin()
	return &Repository{db: db}, db.Error
}

func (r *Repository) Rollback() error {
	return r.db.Rollback().Error
}

func (r *Repository) Commit() error {
	return r.db.Commit().Error
}

func NewRepository() (*Repository, error) {
	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	username := os.Getenv("PG_USERNAME")
	password := os.Getenv("PG_PASSWORD")
	database := os.Getenv("PG_DATABASE")
	sslmode := os.Getenv("PG_SSLMODE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, username, password, database, sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Repository{db: db}, nil
}
