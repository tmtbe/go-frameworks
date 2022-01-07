package application

import (
	"github.com/gin-gonic/gin"
	user2 "test/internal/gen/restapi/operations/user"
	"test/internal/pkg/app"
)

type UserApplicationImpl struct {
}

func NewUserApplicationImpl() *UserApplicationImpl {
	return &UserApplicationImpl{}
}

func (u UserApplicationImpl) CreateUser(ctx *gin.Context, params *user2.CreateUserParams) *app.Response {
	//TODO implement me
	panic("implement me")
}

func (u UserApplicationImpl) CreateUsersWithArrayInput(ctx *gin.Context, params *user2.CreateUsersWithArrayInputParams) *app.Response {
	//TODO implement me
	panic("implement me")
}

func (u UserApplicationImpl) CreateUsersWithListInput(ctx *gin.Context, params *user2.CreateUsersWithListInputParams) *app.Response {
	//TODO implement me
	panic("implement me")
}

func (u UserApplicationImpl) DeleteUser(ctx *gin.Context, params *user2.DeleteUserParams) *app.Response {
	//TODO implement me
	panic("implement me")
}

func (u UserApplicationImpl) GetUserByName(ctx *gin.Context, params *user2.GetUserByNameParams) *app.Response {
	//TODO implement me
	panic("implement me")
}

func (u UserApplicationImpl) LoginUser(ctx *gin.Context, params *user2.LoginUserParams) *app.Response {
	//TODO implement me
	panic("implement me")
}

func (u UserApplicationImpl) LogoutUser(ctx *gin.Context) *app.Response {
	//TODO implement me
	panic("implement me")
}

func (u UserApplicationImpl) UpdateUser(ctx *gin.Context, params *user2.UpdateUserParams) *app.Response {
	//TODO implement me
	panic("implement me")
}
