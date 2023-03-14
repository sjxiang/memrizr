package model

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// type UserService interface {
// 	Get(ctx context.Context, uid uuid.UUID) (*User, error)
// 	Signup(ctx context.Context, u *User) error
// }

// type TokenService interface {
// 	NewPairFromUser(ctx context.Context, u *User, prevTokenID string) (*TokenPair, error)
// }


type UserRepository interface {
	FindByID(ctx context.Context, uid uuid.UUID) (*User, error)
}

type UserRepositoryImpl struct {
	logger *zap.SugaredLogger
	db     *gorm.DB
}


func NewUserRepositoryImpl(logger *zap.SugaredLogger, db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		logger: logger,
		db:     db,
	}	
}


func (impl *UserRepositoryImpl) FindByID(ctx context.Context, uid uuid.UUID) (*User, error) {
	panic("Method not implemented")
}