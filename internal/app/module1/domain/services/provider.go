package services

import "github.com/google/wire"

// ProviderSet 注册所有Services
var ProviderSet = wire.NewSet(
	NewUserDetailServiceImpl,
	wire.Bind(new(UserDetailService), new(*UserDetailServiceImpl)),
)
