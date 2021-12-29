package application

import (
	"go.uber.org/zap"
	"test/internal/app/module1/domain/services"
	"test/internal/pkg/app"
)

type UserDetailApplication struct {
	logger  *zap.Logger
	service services.UserDetailService
}

func NewDetailsApplication(c *app.Context, logger *zap.Logger, s services.UserDetailService) *UserDetailApplication {
	u := &UserDetailApplication{
		logger:  logger,
		service: s,
	}
	c.Add("user_detail_application", u)
	return u
}

func (da *UserDetailApplication) GetUserDetail(id uint64) (*services.UserDetail, error) {
	return da.service.GetUserDetail(id)
}
