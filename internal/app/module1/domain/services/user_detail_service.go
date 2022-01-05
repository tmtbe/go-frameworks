package services

import (
	"go.uber.org/zap"
	expose2 "test/internal/app/module1/application/services"
	"test/internal/app/module1/domain/repos"
	"test/internal/pkg/app"
)

type UserDetailServiceImpl struct {
	logger           *zap.Logger
	detailRepository repos.DetailRepository
	userRepository   repos.UserRepository
}

func NewUserDetailServiceImpl(
	logger *zap.Logger,
	detailRepository repos.DetailRepository,
	userRepository repos.UserRepository,
) *UserDetailServiceImpl {
	u := &UserDetailServiceImpl{
		logger:           logger.With(zap.String("type", "UserDetailServiceImpl")),
		detailRepository: detailRepository,
		userRepository:   userRepository,
	}
	return u
}

func (s *UserDetailServiceImpl) GetUserDetail(ID uint64) (p *expose2.UserDetail, err error) {
	d := s.detailRepository.FindDetailById(ID)
	u := s.userRepository.FindUserById(ID)
	p = fromInfraDetailAndUser(d, u)
	if p == nil {
		err = app.BusinessError("Can't find any user or detail")
	}
	return
}

func fromInfraDetailAndUser(
	detail *repos.DetailRecord,
	user *repos.UserRecord,
) *expose2.UserDetail {
	if detail == nil {
		return nil
	}
	if user == nil {
		return nil
	}
	return &expose2.UserDetail{
		ID:          detail.ID,
		UserName:    user.UserName,
		Password:    user.Password,
		Email:       user.Email,
		Name:        detail.Name,
		Price:       detail.Price,
		CreatedTime: detail.CreatedAt,
	}
}
