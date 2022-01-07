package usercase

import (
	"go.uber.org/zap"
	"test/internal/app/module1/domain/services"
)

type UserDetailUsercase struct {
	logger  *zap.Logger
	service services.UserDetailService
}

func NewUserDetailsUsercase(logger *zap.Logger, s services.UserDetailService) *UserDetailUsercase {
	u := &UserDetailUsercase{
		logger:  logger,
		service: s,
	}
	return u
}

func (da *UserDetailUsercase) GetUserDetail(id uint64) (*services.UserDetail, error) {
	return da.service.GetUserDetail(id)
}
