package repository

import (
	"context"

	"go.uber.org/zap"
	
	"github.com/sjxiang/memrizr/account/model"
)

type RedisTokenRepository interface {
	NewPairFromUser(ctx context.Context, u *model.User, prevTokenID string) (*model.TokenPair, error)
}

type TokenRepositoryImpl struct {
	logger *zap.SugaredLogger
}

func NewTokenRepositoryImpl(_logger *zap.SugaredLogger) *TokenRepositoryImpl {
	return &TokenRepositoryImpl{
		logger: _logger,
	}
}

func (impl *TokenRepositoryImpl) NewPairFromUser(ctx context.Context, u *model.User, prevTokenID string) (*model.TokenPair, error) {
	panic("Method not implemented")
}
