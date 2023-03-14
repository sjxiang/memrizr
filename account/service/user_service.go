package service

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/sjxiang/memrizr/account/model"
)


type UserService interface {
	Get(ctx context.Context, uid uuid.UUID) (*model.User, error)
	Signup(ctx context.Context, u *model.User) error
}

type UserServiceImpl struct {
	logger    *zap.SugaredLogger
	userRepo  model.UserRepository
}

func NewUserServiceImpl(logger *zap.SugaredLogger, userRepo model.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		logger:   logger,
		userRepo: userRepo,
	}
}


func (impl *UserServiceImpl) Get(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	u, err := impl.userRepo.FindByID(ctx, uid)
	return u, err
}

func (impl *UserServiceImpl) Signup(ctx context.Context, u *model.User) error {
	panic("Method not implemented")
}