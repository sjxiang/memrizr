package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/sjxiang/memrizr/account/model"
)


type PGUserRepository interface {
	FindByID(ctx context.Context, uid uuid.UUID) (*model.User, error)
	Create(ctx context.Context, u *model.User) error
}

type UserRepositoryImpl struct {
	logger *zap.SugaredLogger
	db     *gorm.DB
	DB     *sqlx.DB
}


func NewUserRepositoryImpl(logger *zap.SugaredLogger, DB *sqlx.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		logger: logger,
		DB:     DB,
	}	
}


func (impl *UserRepositoryImpl) FindByID(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	panic("Method not implemented")
}


func (impl *UserRepositoryImpl) Create(ctx context.Context, u *model.User) error {
	panic("Method not implemented")
}
