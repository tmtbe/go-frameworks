package application

import (
	"go.uber.org/zap"
	"test/internal/app/module1/domain/expose"
)

type UserDetailApplication struct {
	logger  *zap.Logger
	service expose.UserDetailService
}

func NewUserDetailsApplication(logger *zap.Logger, s expose.UserDetailService) *UserDetailApplication {
	u := &UserDetailApplication{
		logger:  logger,
		service: s,
	}
	return u
}

func (da *UserDetailApplication) GetUserDetail(id uint64) (*expose.UserDetail, error) {
	return da.service.GetUserDetail(id)
}
