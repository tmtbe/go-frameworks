package application

import (
	"go.uber.org/zap"
	"test/internal/app/module1/domain/services"
)

type UserDetailApplication struct {
	logger  *zap.Logger
	service services.UserDetailService
}

func NewDetailsApplication(logger *zap.Logger, s services.UserDetailService) *UserDetailApplication {
	return &UserDetailApplication{
		logger:  logger,
		service: s,
	}
}

func (da *UserDetailApplication) GetUserDetail(id uint64) (*services.UserDetail, error) {
	return da.service.GetUserDetail(id)
}
