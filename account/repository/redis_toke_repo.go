package repository

import (
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"
	
	"github.com/sjxiang/memrizr/account/model"
)

type RedisTokenRepository interface {
	NewPairFromUser(ctx context.Context, u *model.User, prevTokenID string) (*model.TokenPair, error)
}

type TokenRepositoryImpl struct {
	logger *zap.SugaredLogger
	db     *gorm.DB
}

func NewTokenRepositoryImpl(logger *zap.SugaredLogger, db *gorm.DB) *TokenRepositoryImpl {
	return &TokenRepositoryImpl{
		logger: logger,
		db:     db,	
	}
}

func (impl *TokenRepositoryImpl) NewPairFromUser(ctx context.Context, u *model.User, prevTokenID string) (*model.TokenPair, error) {
	panic("Method not implemented")
}
