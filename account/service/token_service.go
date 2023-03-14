package service

import (
	"context"

	"go.uber.org/zap"

	"github.com/sjxiang/memrizr/account/model"
	repo "github.com/sjxiang/memrizr/account/repository"
)


type TokenService interface {
	NewPairFromUser(ctx context.Context, u *model.User, prevTokenID string) (*model.TokenPair, error)	
	// Signout
	// ValidateRefreshToken
	// ValidateIDToken
}

type TokenServiceImpl struct {
	logger    *zap.SugaredLogger
	tokenRepo  repo.RedisTokenRepository
}

func NewTokenServiceImpl(_logger *zap.SugaredLogger, tokenRepo repo.RedisTokenRepository) *TokenServiceImpl {
	return &TokenServiceImpl{
		logger:    _logger,
		tokenRepo: tokenRepo,
	}
}

func (impl *TokenServiceImpl) NewPairFromUser(ctx context.Context, u *model.User, prevTokenID string) (*model.TokenPair, error) {
	panic("TODO")	
}

