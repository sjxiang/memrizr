package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"go.uber.org/zap"

	"github.com/sjxiang/memrizr/account/model"
	"github.com/sjxiang/memrizr/account/util/apperrors"
)


type PGUserRepository interface {
	FindByID(ctx context.Context, uid uuid.UUID) (*model.User, error)
	Create(ctx context.Context, u *model.User) error
}

type UserRepositoryImpl struct {
	logger *zap.SugaredLogger
	DB     *sqlx.DB
}


func NewUserRepositoryImpl(_logger *zap.SugaredLogger, db *sqlx.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		logger: _logger,
		DB:     db,
	}	
}


func (impl *UserRepositoryImpl) FindByID(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	user := &model.User{}

	query := "SELECT * FROM users WHERE uid=$1"

	if err := impl.DB.Get(user, query, uid); err != nil {
		return user, apperrors.NewNotFound("uid", uid.String())
	}

	return user, nil 
}


func (impl *UserRepositoryImpl) Create(ctx context.Context, u *model.User) error {
	query := "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING *"

	if err := impl.DB.Get(u, query, u.Email, u.Password); err != nil {
		// 检查唯一约束
		if err, ok := err.(*pq.Error); ok && err.Code.Name() == "unique_violation" {
			impl.logger.Errorf("Could not create a user with email: %v. Reason: %v\n", u.Email, err.Code.Name())
			return apperrors.NewConflict("email", u.Email)
		}

		impl.logger.Errorf("Could not create a user with email: %v. Reason: %v\n", u.Email, err)
		return apperrors.NewInternal()
	}

	return nil
}
