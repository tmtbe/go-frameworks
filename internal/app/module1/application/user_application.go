package application

import (
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"test/internal/gen/models"
	"test/internal/gen/restapi/operations/user"
	"test/internal/pkg/app"
)

type UserApplicationImpl struct {
}

func (u UserApplicationImpl) GetUsersUserID(ctx *gin.Context, params *user.GetUsersUserIDParams) *app.Response {
	return &app.Response{
		Code: 200,
		Body: &models.User{
			CreateDate:    strfmt.Date{},
			DateOfBirth:   nil,
			Email:         nil,
			EmailVerified: nil,
			FirstName:     nil,
			ID:            0,
			LastName:      nil,
		},
	}
}

func (u UserApplicationImpl) PostUser(ctx *gin.Context, params *user.PostUserParams) *app.Response {
	//TODO implement me
	panic("implement me")
}

func NewUserApplicationImpl() *UserApplicationImpl {
	return &UserApplicationImpl{}
}
