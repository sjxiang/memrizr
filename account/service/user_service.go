package service

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/sjxiang/memrizr/account/model"
	repo "github.com/sjxiang/memrizr/account/repository"
	"github.com/sjxiang/memrizr/account/util"
	"github.com/sjxiang/memrizr/account/util/apperrors"
)


type UserService interface {
	Get(ctx context.Context, uid uuid.UUID) (*model.User, error)
	Signup(ctx context.Context, u *model.User) error
	// Signin
	// UpdateDeatils
	// SetProfileImage
	// ClearProfileImage, Delete
}

type UserServiceImpl struct {
	logger    *zap.SugaredLogger
	userRepo  repo.PGUserRepository
}

func NewUserServiceImpl(logger *zap.SugaredLogger, userRepo repo.PGUserRepository) *UserServiceImpl {
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
	pw, err := util.HashPassword(u.Password)
	
	if err != nil {
		impl.logger.Errorf("Unable to signup user for email: %v\n", u.Email)
		return apperrors.NewInternal()
	}

	u.Password = pw

	if err := impl.userRepo.Create(ctx, u); err != nil {
		return err
	}

	return nil
}

	// panic("Method not implemented")